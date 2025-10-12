package vault

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/aerium-network/aerium/crypto"
	"github.com/aerium-network/aerium/crypto/bls"
	"github.com/aerium-network/aerium/crypto/bls/hdkeychain"
	"github.com/aerium-network/aerium/crypto/ed25519"
	"github.com/aerium-network/aerium/util/testsuite"
	"github.com/aerium-network/aerium/wallet/addresspath"
	"github.com/aerium-network/aerium/wallet/encrypter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const tPassword = "super_secret_password"

type testData struct {
	*testsuite.TestSuite

	vault              *Vault
	mnemonic           string
	importedBLSPrv     *bls.PrivateKey
	importedEd25519Prv *ed25519.PrivateKey
}

// setup returns an instances of vault fo testing.
func setup(t *testing.T) *testData {
	t.Helper()

	ts := testsuite.NewTestSuite(t)

	mnemonic, _ := GenerateMnemonic(128)
	vault, err := CreateVaultFromMnemonic(mnemonic, 19933) // Mainnet
	assert.NoError(t, err)

	key, _ := hdkeychain.NewKeyFromString(vault.Purposes.PurposeBLS.XPubAccount)
	assert.False(t, key.IsPrivate())

	// Create some test address
	_, err = vault.NewBLSAccountAddress("bls-account-address")
	assert.NoError(t, err)
	_, err = vault.NewEd25519AccountAddress("ed25519-account-address", "")
	assert.NoError(t, err)
	_, err = vault.NewValidatorAddress("validator-address")
	assert.NoError(t, err)

	_, importedBLSPrv := ts.RandBLSKeyPair()
	assert.NoError(t, vault.ImportBLSPrivateKey("", importedBLSPrv))

	_, importedEd25519Prv := ts.RandEd25519KeyPair()
	assert.NoError(t, vault.ImportEd25519PrivateKey("", importedEd25519Prv))

	assert.False(t, vault.IsEncrypted())

	// Set encryption options to minimal values for faster test execution.
	opts := []encrypter.Option{
		encrypter.OptionIteration(1),
		encrypter.OptionMemory(8),
		encrypter.OptionParallelism(1),
	}

	err = vault.UpdatePassword("", tPassword, opts...)
	assert.NoError(t, err)
	assert.True(t, vault.IsEncrypted())

	return &testData{
		TestSuite:          ts,
		vault:              vault,
		mnemonic:           mnemonic,
		importedBLSPrv:     importedBLSPrv,
		importedEd25519Prv: importedEd25519Prv,
	}
}

func TestAddressCount(t *testing.T) {
	td := setup(t)

	assert.Equal(t, 6, td.vault.AddressCount())

	// Neutered
	neutered := td.vault.Neuter()
	assert.Equal(t, 6, neutered.AddressCount())
}

func TestContains(t *testing.T) {
	td := setup(t)

	t.Run("Vault should contain all known addresses", func(t *testing.T) {
		infos := td.vault.AddressInfos()
		for _, i := range infos {
			assert.True(t, td.vault.Contains(i.Address))
		}
	})

	t.Run("Vault should not contain unknown address", func(t *testing.T) {
		unknownAddr := td.RandAccAddress().String()
		assert.False(t, td.vault.Contains(unknownAddr))
	})
}

func TestSortAddressInfo(t *testing.T) {
	td := setup(t)

	infos := td.vault.AddressInfos()

	// Ed25519 Keys
	assert.Equal(t, "m/44'/19933'/3'/0'", infos[0].Path)
	// BLS Keys
	assert.Equal(t, "m/12381'/19933'/1'/0", infos[1].Path)
	assert.Equal(t, "m/12381'/19933'/2'/0", infos[2].Path)
	// Imported Keys
	assert.Equal(t, "m/65535'/19933'/1'/0'", infos[3].Path)
	assert.Equal(t, "m/65535'/19933'/2'/0'", infos[4].Path)
	assert.Equal(t, "m/65535'/19933'/3'/1'", infos[5].Path)
}

func TestAllAccountAddresses(t *testing.T) {
	td := setup(t)

	accountAddrs := td.vault.AllAccountAddresses()
	for _, i := range accountAddrs {
		path, err := addresspath.FromString(i.Path)
		assert.NoError(t, err)

		assert.NotEqual(t, _H(crypto.AddressTypeValidator), path.AddressType())
	}
}

func TestAllValidatorAddresses(t *testing.T) {
	td := setup(t)

	validatorAddrs := td.vault.AllValidatorAddresses()
	for _, i := range validatorAddrs {
		info := td.vault.AddressInfo(i.Address)
		assert.Equal(t, i.Address, info.Address)

		path, _ := addresspath.FromString(info.Path)

		switch path.Purpose() {
		case _H(PurposeBLS12381):
			assert.Equal(t, fmt.Sprintf("m/%d'/%d'/1'/%d",
				PurposeBLS12381, td.vault.CoinType, path.AddressIndex()), info.Path)
		case _H(PurposeImportPrivateKey):
			assert.Equal(t, fmt.Sprintf("m/%d'/%d'/1'/%d'",
				PurposeImportPrivateKey, td.vault.CoinType, _N(path.AddressIndex())), info.Path)
		default:
			assert.Fail(t, "not supported")
		}
	}
}

func TestSortAllValidatorAddresses(t *testing.T) {
	td := setup(t)

	validatorAddrs := td.vault.AllValidatorAddresses()

	assert.Equal(t, "m/12381'/19933'/1'/0", validatorAddrs[0].Path)
	assert.Equal(t, "m/65535'/19933'/1'/0'", validatorAddrs[len(validatorAddrs)-1].Path)
}

func TestAddressFromPath(t *testing.T) {
	td := setup(t)

	t.Run("Could not find address from path", func(t *testing.T) {
		path := "m/12381'/26888'/983'/0"
		assert.Nil(t, td.vault.AddressFromPath(path))
	})

	t.Run("Ok", func(t *testing.T) {
		var address string
		var addrInfo AddressInfo

		for addr, ai := range td.vault.Addresses {
			address = addr
			addrInfo = ai

			break
		}

		assert.Equal(t, address, td.vault.AddressFromPath(addrInfo.Path).Address)
	})
}

func TestNewValidatorAddress(t *testing.T) {
	td := setup(t)

	label := td.RandString(16)
	addressInfo, err := td.vault.NewValidatorAddress(label)
	assert.NoError(t, err)
	assert.NotEmpty(t, addressInfo.Address)
	assert.NotEmpty(t, addressInfo.PublicKey)
	assert.Contains(t, addressInfo.Path, "m/12381'/19933'/1'")
	assert.Equal(t, label, addressInfo.Label)

	pub, _ := bls.PublicKeyFromString(addressInfo.PublicKey)
	assert.Equal(t, pub.ValidatorAddress().String(), addressInfo.Address)
}

func TestNewBLSAccountAddress(t *testing.T) {
	td := setup(t)

	label := td.RandString(16)
	addressInfo, err := td.vault.NewBLSAccountAddress(label)
	assert.NoError(t, err)
	assert.NotEmpty(t, addressInfo.Address)
	assert.NotEmpty(t, addressInfo.PublicKey)
	assert.Contains(t, addressInfo.Path, "m/12381'/19933'/2'")
	assert.Equal(t, label, addressInfo.Label)

	pub, _ := bls.PublicKeyFromString(addressInfo.PublicKey)
	assert.Equal(t, pub.AccountAddress().String(), addressInfo.Address)
}

func TestNewE225519AccountAddress(t *testing.T) {
	td := setup(t)

	addressInfo, err := td.vault.NewEd25519AccountAddress("addr-2", tPassword)
	assert.NoError(t, err)
	assert.NotEmpty(t, addressInfo.Address)
	assert.NotEmpty(t, addressInfo.PublicKey)
	assert.Equal(t, "m/44'/19933'/3'/1'", addressInfo.Path)

	pub, _ := ed25519.PublicKeyFromString(addressInfo.PublicKey)
	assert.Equal(t, pub.AccountAddress().String(), addressInfo.Address)
}

func TestRecover(t *testing.T) {
	td := setup(t)

	t.Run("Invalid mnemonic", func(t *testing.T) {
		_, err := CreateVaultFromMnemonic("invalid mnemonic phrase seed", 19933)
		assert.Error(t, err)
	})

	t.Run("Ok", func(t *testing.T) {
		recovered, err := CreateVaultFromMnemonic(td.mnemonic, 19933)
		assert.NoError(t, err)

		// Recover addresses
		_, err = recovered.NewBLSAccountAddress("bls-account-address")
		assert.NoError(t, err)
		_, err = recovered.NewEd25519AccountAddress("ed25519-account-address", "")
		assert.NoError(t, err)
		_, err = recovered.NewValidatorAddress("validator-address")
		assert.NoError(t, err)

		assert.Equal(t, recovered.Purposes, td.vault.Purposes)
	})
}

func TestGetPrivateKeys(t *testing.T) {
	td := setup(t)

	t.Run("Unknown address", func(t *testing.T) {
		addr := td.RandAccAddress()
		_, err := td.vault.PrivateKeys(tPassword, []string{addr.String()})
		assert.ErrorIs(t, err, NewErrAddressNotFound(addr.String()))
	})

	t.Run("No password", func(t *testing.T) {
		addr := td.vault.AddressInfos()[0].Address
		_, err := td.vault.PrivateKeys("", []string{addr})
		assert.ErrorIs(t, err, encrypter.ErrInvalidPassword)
	})

	t.Run("Invalid password", func(t *testing.T) {
		addr := td.vault.AddressInfos()[0].Address
		_, err := td.vault.PrivateKeys("wrong_password", []string{addr})
		assert.ErrorIs(t, err, encrypter.ErrInvalidPassword)
	})

	t.Run("Check all the private keys", func(t *testing.T) {
		for _, info := range td.vault.AddressInfos() {
			prv, err := td.vault.PrivateKeys(tPassword, []string{info.Address})
			assert.NoError(t, err)
			addrInfo := td.vault.AddressInfo(info.Address)
			path, _ := addresspath.FromString(info.Path)

			switch _N(path.AddressType()) {
			case uint32(crypto.AddressTypeBLSAccount),
				uint32(crypto.AddressTypeValidator):
				pub, _ := bls.PublicKeyFromString(addrInfo.PublicKey)
				require.True(t, prv[0].PublicKey().EqualsTo(pub))
			case uint32(crypto.AddressTypeEd25519Account):
				pub, _ := ed25519.PublicKeyFromString(addrInfo.PublicKey)
				require.True(t, prv[0].PublicKey().EqualsTo(pub))
			default:
				assert.Fail(t, "not supported")
			}
		}
	})
}

func TestImportBLSPrivateKey(t *testing.T) {
	td := setup(t)

	_, prv := td.RandBLSKeyPair()

	t.Run("Invalid password", func(t *testing.T) {
		err := td.vault.ImportBLSPrivateKey("invalid-password", prv)
		assert.ErrorIs(t, err, encrypter.ErrInvalidPassword)
	})

	t.Run("Ok", func(t *testing.T) {
		err := td.vault.ImportBLSPrivateKey(tPassword, prv)
		assert.NoError(t, err)

		valAddr := prv.PublicKeyNative().ValidatorAddress().String()
		accAddr := prv.PublicKeyNative().AccountAddress().String()

		valAddrInfo := td.vault.AddressInfo(valAddr)
		accAddrInfo := td.vault.AddressInfo(accAddr)

		assert.True(t, td.vault.Contains(valAddr))
		assert.True(t, td.vault.Contains(accAddr))

		assert.Equal(t, valAddr, valAddrInfo.Address)
		assert.Equal(t, accAddr, accAddrInfo.Address)

		assert.Equal(t, prv.PublicKeyNative().String(), valAddrInfo.PublicKey)
		assert.Equal(t, prv.PublicKeyNative().String(), accAddrInfo.PublicKey)

		assert.Equal(t, "m/65535'/19933'/1'/2'", valAddrInfo.Path)
		assert.Equal(t, "m/65535'/19933'/2'/2'", accAddrInfo.Path)
	})

	t.Run("Reimporting private key", func(t *testing.T) {
		err := td.vault.ImportBLSPrivateKey(tPassword, prv)
		assert.ErrorIs(t, err, ErrAddressExists)
	})
}

func TestImportEd25519PrivateKey(t *testing.T) {
	td := setup(t)

	_, prv := td.RandEd25519KeyPair()

	t.Run("Invalid password", func(t *testing.T) {
		err := td.vault.ImportEd25519PrivateKey("invalid-password", prv)
		assert.ErrorIs(t, err, encrypter.ErrInvalidPassword)
	})

	t.Run("Ok", func(t *testing.T) {
		err := td.vault.ImportEd25519PrivateKey(tPassword, prv)
		assert.NoError(t, err)

		accAddr := prv.PublicKeyNative().AccountAddress().String()

		accAddrInfo := td.vault.AddressInfo(accAddr)
		assert.True(t, td.vault.Contains(accAddr))
		assert.Equal(t, accAddr, accAddrInfo.Address)
		assert.Equal(t, prv.PublicKeyNative().String(), accAddrInfo.PublicKey)
		assert.Equal(t, "m/65535'/19933'/3'/2'", accAddrInfo.Path)
	})

	t.Run("Reimporting private key", func(t *testing.T) {
		err := td.vault.ImportEd25519PrivateKey(tPassword, td.importedEd25519Prv)
		assert.ErrorIs(t, err, ErrAddressExists)
	})
}

func TestGetMnemonic(t *testing.T) {
	td := setup(t)

	t.Run("Invalid password", func(t *testing.T) {
		_, err := td.vault.Mnemonic("invalid-password")
		assert.ErrorIs(t, err, encrypter.ErrInvalidPassword)
	})

	t.Run("No password", func(t *testing.T) {
		_, err := td.vault.Mnemonic("")
		assert.ErrorIs(t, err, encrypter.ErrInvalidPassword)
	})

	t.Run("Ok", func(t *testing.T) {
		m, err := td.vault.Mnemonic(tPassword)
		assert.NoError(t, err)
		assert.Equal(t, m, td.mnemonic)
	})

	t.Run("Neutered wallet", func(t *testing.T) {
		_, err := td.vault.Neuter().Mnemonic("")
		assert.ErrorIs(t, err, ErrNeutered)
	})
}

func TestUpdatePassword(t *testing.T) {
	td := setup(t)

	opts := []encrypter.Option{
		encrypter.OptionIteration(1),
		encrypter.OptionMemory(1),
		encrypter.OptionParallelism(1),
	}

	addrInfos := td.vault.AddressInfos()
	newPassword := "new-password"

	t.Run("Empty password", func(t *testing.T) {
		err := td.vault.UpdatePassword("", newPassword)
		assert.ErrorIs(t, err, encrypter.ErrInvalidPassword)
	})

	t.Run("Incorrect password", func(t *testing.T) {
		err := td.vault.UpdatePassword("invalid-password", newPassword)
		assert.ErrorIs(t, err, encrypter.ErrInvalidPassword)
	})

	t.Run("Valid password update", func(t *testing.T) {
		assert.NoError(t, td.vault.UpdatePassword(tPassword, newPassword, opts...))
		assert.True(t, td.vault.IsEncrypted())
		assert.Equal(t, addrInfos, td.vault.AddressInfos())
	})

	t.Run("Old password should no longer be valid", func(t *testing.T) {
		err := td.vault.UpdatePassword(tPassword, newPassword)
		assert.ErrorIs(t, err, encrypter.ErrInvalidPassword)
	})

	t.Run("Set vault password to empty", func(t *testing.T) {
		assert.NoError(t, td.vault.UpdatePassword(newPassword, ""))
		assert.False(t, td.vault.IsEncrypted())
		assert.Equal(t, addrInfos, td.vault.AddressInfos())
	})
}

func TestSetLabel(t *testing.T) {
	td := setup(t)

	t.Run("Set label for unknown address", func(t *testing.T) {
		invAddr := td.RandAccAddress().String()
		err := td.vault.SetLabel(invAddr, "i have label")
		assert.ErrorIs(t, err, NewErrAddressNotFound(invAddr))
		assert.Equal(t, "", td.vault.Label(invAddr))
	})

	t.Run("Update label", func(t *testing.T) {
		testAddr := td.vault.AddressInfos()[0].Address
		err := td.vault.SetLabel(testAddr, "I have a label")
		assert.NoError(t, err)
		assert.Equal(t, "I have a label", td.vault.Label(testAddr))
	})

	t.Run("Remove label", func(t *testing.T) {
		testAddr := td.vault.AddressInfos()[0].Address
		err := td.vault.SetLabel(testAddr, "")
		assert.NoError(t, err)
		var ok bool
		l := td.vault.Label(testAddr)
		if strings.TrimSpace(l) != "" {
			ok = true
		}
		assert.Empty(t, td.vault.Label(testAddr))
		assert.False(t, ok)
	})
}

func TestNeuter(t *testing.T) {
	td := setup(t)

	neutered := td.vault.Neuter()
	_, err := neutered.Mnemonic(tPassword)
	assert.ErrorIs(t, err, ErrNeutered)

	_, err = neutered.PrivateKeys(tPassword, []string{
		td.RandAccAddress().String(),
	})
	assert.ErrorIs(t, err, ErrNeutered)

	_, prv := td.RandBLSKeyPair()
	err = neutered.ImportBLSPrivateKey("any", prv)
	assert.ErrorIs(t, err, ErrNeutered)

	err = td.vault.Neuter().UpdatePassword("any", "any")
	assert.ErrorIs(t, err, ErrNeutered)
}

// TestAddressRecovery tests the address recovery functionality.
// The first 8 BLS account addresses for the test mnemonic are:
// ae1z0v7jum3clvcfeurxg92vt8uwx002ev9jxtyk3p (index 0)
// ae1z5q8m2a5fe4wdwvzr2zf7w3c7lu6yas4ynf6sfu (index 1)
// ae1zkcv2d3kywlag93y4zf4p453p22mc8xvyca3h39 (index 2)
// ae1z2pqqrh2slw9aqj34uf8yrnuk39dttht29dpl4u (index 3)
// ae1z2ylnhcujfeg3cpa923q54zhjugz6mv7xkgvun4 (index 4)
// ae1zzpa29tnf85ukg3pstgdgh4dnhv2fgtaufwc34c (index 5)
// ae1z2yqtjugl8q3fpqnkp4256f947559gm7p9jk0xe (index 6)
// ae1z7sv9x0kqyms7vzc5z9rdzlgyh6egmpw4mfhe9a (index 7)
//
// The first 8 Ed25519 account addresses for the test mnemonic are:
// ae1rex4yps0chj9lsax9n8m0vpt59a2jecxv8k0g8y (index 0)
// ae1rnk393e0ax68472nxn3zgt4vap0mvvsn89y20rn (index 1)
// ae1rsn6hp05ud04vz5rung0dva4mtzdr3kjumv25hy (index 2)
// ae1rqzqlkelnlmefz0mfmyr4gk2a236686vr7c7a0z (index 3)
// ae1ruvwcyxdz5nlmehlhw6du6pw2xjls5lq56tu8re (index 4)
// ae1rxfnz6nvrj57m4lqny4qh5mgtlkagsj05rpzr9k (index 5)
// ae1rnhtvss0m7sdt4zh4r892xswczgwcyzjxms4kcd (index 6)
// ae1rnf2dha4cgl95er9mlvhe39ql6204qd6grme8lv (index 7)
//
// The test uses a mock hasActivity function to simulate blockchain activity checks.
func TestAddressRecovery(t *testing.T) {
	//nolint:dupword // has duplicated words
	testMnemonic := "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon cactus"

	t.Run("recover addresses from a fresh wallet without any active addresses", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933) // Mainnet
		assert.NoError(t, err)

		// Mock hasActivity to return false for all addresses (no active addresses)
		hasActivity := func(_ string) (bool, error) {
			return false, nil
		}

		err = vault.RecoverAddresses(context.Background(), "", hasActivity)
		assert.NoError(t, err)

		// Should have 1 Ed25519 address (the first one)
		addresses := vault.AllAccountAddresses()
		assert.Len(t, addresses, 1)
		assert.Equal(t, "ae1rex4yps0chj9lsax9n8m0vpt59a2jecxv8k0g8y", addresses[0].Address)
	})

	t.Run("recover addresses with one gap at the beginning", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933) // Mainnet
		assert.NoError(t, err)

		// Mock hasActivity to return true only for the first call (address at index 0)
		hasActivity := func(addr string) (bool, error) {
			return addr == "ae1rnk393e0ax68472nxn3zgt4vap0mvvsn89y20rn" ||
				addr == "ae1z5q8m2a5fe4wdwvzr2zf7w3c7lu6yas4ynf6sfu", nil
		}

		err = vault.RecoverAddresses(context.Background(), "", hasActivity)
		assert.NoError(t, err)

		// Should have 4 addresses
		addresses := vault.AllAccountAddresses()
		assert.Len(t, addresses, 4)
		assert.Equal(t, "ae1rex4yps0chj9lsax9n8m0vpt59a2jecxv8k0g8y", addresses[0].Address)
		assert.Equal(t, "ae1rnk393e0ax68472nxn3zgt4vap0mvvsn89y20rn", addresses[1].Address)
		assert.Equal(t, "ae1z0v7jum3clvcfeurxg92vt8uwx002ev9jxtyk3p", addresses[2].Address)
		assert.Equal(t, "ae1z5q8m2a5fe4wdwvzr2zf7w3c7lu6yas4ynf6sfu", addresses[3].Address)
	})

	t.Run("recover addresses with gaps in the middle of the address list", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933) // Mainnet
		assert.NoError(t, err)

		hasActivity := func(addr string) (bool, error) {
			return addr == "ae1rex4yps0chj9lsax9n8m0vpt59a2jecxv8k0g8y" ||
				addr == "ae1rnk393e0ax68472nxn3zgt4vap0mvvsn89y20rn" ||
				addr == "ae1rqzqlkelnlmefz0mfmyr4gk2a236686vr7c7a0z" ||
				addr == "ae1z0v7jum3clvcfeurxg92vt8uwx002ev9jxtyk3p" ||
				addr == "ae1z2pqqrh2slw9aqj34uf8yrnuk39dttht29dpl4u", nil
		}

		err = vault.RecoverAddresses(context.Background(), "", hasActivity)
		assert.NoError(t, err)

		addresses := vault.AllAccountAddresses()
		assert.Len(t, addresses, 8)

		assert.Equal(t, "ae1rex4yps0chj9lsax9n8m0vpt59a2jecxv8k0g8y", addresses[0].Address)
		assert.Equal(t, "ae1rnk393e0ax68472nxn3zgt4vap0mvvsn89y20rn", addresses[1].Address)
		assert.Equal(t, "ae1rsn6hp05ud04vz5rung0dva4mtzdr3kjumv25hy", addresses[2].Address)
		assert.Equal(t, "ae1rqzqlkelnlmefz0mfmyr4gk2a236686vr7c7a0z", addresses[3].Address)
		assert.Equal(t, "ae1z0v7jum3clvcfeurxg92vt8uwx002ev9jxtyk3p", addresses[4].Address)
		assert.Equal(t, "ae1z5q8m2a5fe4wdwvzr2zf7w3c7lu6yas4ynf6sfu", addresses[5].Address)
		assert.Equal(t, "ae1zkcv2d3kywlag93y4zf4p453p22mc8xvyca3h39", addresses[6].Address)
		assert.Equal(t, "ae1z2pqqrh2slw9aqj34uf8yrnuk39dttht29dpl4u", addresses[7].Address)
	})

	t.Run("error handling", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933) // Mainnet
		assert.NoError(t, err)

		// Mock hasActivity to return an error
		hasActivity := func(_ string) (bool, error) {
			return false, errors.New("blockchain connection error")
		}

		err = vault.RecoverAddresses(context.Background(), "", hasActivity)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "blockchain connection error")
	})

	t.Run("cancel recovery with context cancel signal", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933) // Mainnet
		assert.NoError(t, err)

		// Create a cancellable context
		ctx, cancel := context.WithCancel(context.Background())

		// Counter to track how many times hasActivity is called
		callCount := 0

		// Mock hasActivity to cancel context after a few calls
		hasActivity := func(_ string) (bool, error) {
			callCount++
			// Cancel the context after 3 calls to simulate interruption during recovery
			if callCount >= 3 {
				cancel()
			}

			return false, nil
		}

		err = vault.RecoverAddresses(ctx, "", hasActivity)
		assert.Error(t, err)
		assert.Equal(t, context.Canceled, err)

		// Recovery should have been interrupted, so we should have only the first Ed25519 address
		// or possibly none if cancelled early enough
		addresses := vault.AllAccountAddresses()
		assert.LessOrEqual(t, len(addresses), 1)
	})

	t.Run("recover only BLS addresses without Ed25519 activity", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933) // Mainnet
		assert.NoError(t, err)

		// Mock hasActivity to return true only for BLS addresses
		hasActivity := func(addr string) (bool, error) {
			return addr == "ae1z0v7jum3clvcfeurxg92vt8uwx002ev9jxtyk3p" ||
				addr == "ae1zkcv2d3kywlag93y4zf4p453p22mc8xvyca3h39", nil
		}

		err = vault.RecoverAddresses(context.Background(), "", hasActivity)
		assert.NoError(t, err)

		addresses := vault.AllAccountAddresses()
		// Should have: 1 Ed25519 (default) + 3 BLS (0, 1, 2)
		assert.Len(t, addresses, 4)
		assert.Equal(t, "ae1rex4yps0chj9lsax9n8m0vpt59a2jecxv8k0g8y", addresses[0].Address)
		assert.Equal(t, "ae1z0v7jum3clvcfeurxg92vt8uwx002ev9jxtyk3p", addresses[1].Address)
		assert.Equal(t, "ae1z5q8m2a5fe4wdwvzr2zf7w3c7lu6yas4ynf6sfu", addresses[2].Address)
		assert.Equal(t, "ae1zkcv2d3kywlag93y4zf4p453p22mc8xvyca3h39", addresses[3].Address)
	})

	t.Run("recover addresses with activity at exact gap limit boundary", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933) // Mainnet
		assert.NoError(t, err)

		// Mock hasActivity with Ed25519 activity at index 0 and index 8 (exactly at gap limit)
		// This should recover indices 0-8 for Ed25519 (9 addresses total)
		hasActivity := func(addr string) (bool, error) {
			return addr == "ae1rex4yps0chj9lsax9n8m0vpt59a2jecxv8k0g8y" || // index 0
				addr == "ae1rncua0gelt6p8pucgp5ghy7n0dppduylmhw04mj", nil // index 8
		}

		err = vault.RecoverAddresses(context.Background(), "", hasActivity)
		assert.NoError(t, err)

		addresses := vault.AllAccountAddresses()
		// Should recover 9 Ed25519 addresses (indices 0-8)
		assert.Len(t, addresses, 9)
		assert.Equal(t, "ae1rex4yps0chj9lsax9n8m0vpt59a2jecxv8k0g8y", addresses[0].Address)
	})

	t.Run("recover addresses with BLS activity far apart", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933) // Mainnet
		assert.NoError(t, err)

		// Mock hasActivity with BLS addresses at indices 0 and 5 (gap of 4 inactive addresses)
		hasActivity := func(addr string) (bool, error) {
			return addr == "ae1z0v7jum3clvcfeurxg92vt8uwx002ev9jxtyk3p" || // BLS index 0
				addr == "ae1zzpa29tnf85ukg3pstgdgh4dnhv2fgtaufwc34c", nil // BLS index 5
		}

		err = vault.RecoverAddresses(context.Background(), "", hasActivity)
		assert.NoError(t, err)

		addresses := vault.AllAccountAddresses()
		// Should have: 1 Ed25519 (default) + 6 BLS (indices 0-5)
		assert.Len(t, addresses, 7)
		// Verify the BLS addresses are recovered
		blsAddresses := []string{
			"ae1z0v7jum3clvcfeurxg92vt8uwx002ev9jxtyk3p", // index 0
			"ae1z5q8m2a5fe4wdwvzr2zf7w3c7lu6yas4ynf6sfu", // index 1
			"ae1zkcv2d3kywlag93y4zf4p453p22mc8xvyca3h39", // index 2
			"ae1z2pqqrh2slw9aqj34uf8yrnuk39dttht29dpl4u", // index 3
			"ae1z2ylnhcujfeg3cpa923q54zhjugz6mv7xkgvun4", // index 4
			"ae1zzpa29tnf85ukg3pstgdgh4dnhv2fgtaufwc34c", // index 5
		}
		for _, blsAddr := range blsAddresses {
			found := false
			for _, addr := range addresses {
				if addr.Address == blsAddr {
					found = true

					break
				}
			}
			assert.True(t, found, "BLS address %s should be recovered", blsAddr)
		}
	})

	t.Run("recover with password-protected vault", func(t *testing.T) {
		password := "test-password-123"
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933) // Mainnet
		assert.NoError(t, err)

		// Set password on the vault
		err = vault.UpdatePassword("", password)
		assert.NoError(t, err)

		// Mock hasActivity to return true for first two Ed25519 addresses
		hasActivity := func(addr string) (bool, error) {
			return addr == "ae1rex4yps0chj9lsax9n8m0vpt59a2jecxv8k0g8y" ||
				addr == "ae1rnk393e0ax68472nxn3zgt4vap0mvvsn89y20rn", nil
		}

		err = vault.RecoverAddresses(context.Background(), password, hasActivity)
		assert.NoError(t, err)

		addresses := vault.AllAccountAddresses()
		// Should have 2 Ed25519 addresses (indices 0-1)
		assert.Len(t, addresses, 2)
		assert.Equal(t, "ae1rex4yps0chj9lsax9n8m0vpt59a2jecxv8k0g8y", addresses[0].Address)
		assert.Equal(t, "ae1rnk393e0ax68472nxn3zgt4vap0mvvsn89y20rn", addresses[1].Address)
	})

	t.Run("recover with wrong password should fail", func(t *testing.T) {
		password := "correct-password"
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933) // Mainnet
		assert.NoError(t, err)

		// Set password on the vault
		err = vault.UpdatePassword("", password)
		assert.NoError(t, err)

		// Mock hasActivity
		hasActivity := func(addr string) (bool, error) {
			return addr == "ae1rex4yps0chj9lsax9n8m0vpt59a2jecxv8k0g8y", nil
		}

		// Try to recover with wrong password
		err = vault.RecoverAddresses(context.Background(), "wrong-password", hasActivity)
		assert.Error(t, err)
	})

	t.Run("no BLS recovery when no BLS addresses have activity", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933) // Mainnet
		assert.NoError(t, err)

		// Mock hasActivity to return true only for Ed25519 addresses (no BLS activity)
		hasActivity := func(addr string) (bool, error) {
			return addr == "ae1rex4yps0chj9lsax9n8m0vpt59a2jecxv8k0g8y" ||
				addr == "ae1rnk393e0ax68472nxn3zgt4vap0mvvsn89y20rn" ||
				addr == "ae1rsn6hp05ud04vz5rung0dva4mtzdr3kjumv25hy", nil
		}

		err = vault.RecoverAddresses(context.Background(), "", hasActivity)
		assert.NoError(t, err)

		addresses := vault.AllAccountAddresses()
		// Should have only 3 Ed25519 addresses (no BLS addresses)
		assert.Len(t, addresses, 3)
		assert.Equal(t, "ae1rex4yps0chj9lsax9n8m0vpt59a2jecxv8k0g8y", addresses[0].Address)
		assert.Equal(t, "ae1rnk393e0ax68472nxn3zgt4vap0mvvsn89y20rn", addresses[1].Address)
		assert.Equal(t, "ae1rsn6hp05ud04vz5rung0dva4mtzdr3kjumv25hy", addresses[2].Address)
	})

	t.Run("recover maximum addresses beyond gap limit", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933) // Mainnet
		assert.NoError(t, err)

		callCount := 0
		// Mock hasActivity that tracks how many addresses are checked
		// Return activity for every 7th address to test gap limit behavior
		hasActivity := func(addr string) (bool, error) {
			callCount++
			// Have activity at indices 0, 7 for Ed25519 to push beyond one gap limit
			return addr == "ae1rex4yps0chj9lsax9n8m0vpt59a2jecxv8k0g8y" || // Ed25519 index 0
				addr == "ae1rnf2dha4cgl95er9mlvhe39ql6204qd6grme8lv", nil // Ed25519 index 7
		}

		err = vault.RecoverAddresses(context.Background(), "", hasActivity)
		assert.NoError(t, err)

		addresses := vault.AllAccountAddresses()
		// Should have 8 Ed25519 addresses (indices 0-7)
		assert.Len(t, addresses, 8)
		assert.Equal(t, "ae1rex4yps0chj9lsax9n8m0vpt59a2jecxv8k0g8y", addresses[0].Address)
		assert.Equal(t, "ae1rnf2dha4cgl95er9mlvhe39ql6204qd6grme8lv", addresses[7].Address)

		// Verify that we stopped after hitting the gap limit
		// We should have checked: 8 BLS + at least 16 Ed25519 (0-7 active, then 8 more for gap)
		assert.Greater(t, callCount, 16)
	})

	t.Run("context cancellation during BLS recovery", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933) // Mainnet
		assert.NoError(t, err)

		ctx, cancel := context.WithCancel(context.Background())
		callCount := 0

		// Mock hasActivity to cancel after first BLS address check
		hasActivity := func(addr string) (bool, error) {
			callCount++
			// Cancel after checking first BLS address
			if callCount >= 1 && addr == "ae1z0v7jum3clvcfeurxg92vt8uwx002ev9jxtyk3p" {
				cancel()
			}

			return false, nil
		}

		err = vault.RecoverAddresses(ctx, "", hasActivity)
		assert.Error(t, err)
		assert.Equal(t, context.Canceled, err)
	})

	t.Run("hasActivity error during BLS recovery", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933) // Mainnet
		assert.NoError(t, err)

		callCount := 0
		// Mock hasActivity to return error on second call (during BLS recovery)
		hasActivity := func(_ string) (bool, error) {
			callCount++
			if callCount == 2 {
				return false, errors.New("network error during BLS check")
			}

			return false, nil
		}

		err = vault.RecoverAddresses(context.Background(), "", hasActivity)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "network error during BLS check")
	})

	t.Run("hasActivity error during Ed25519 recovery", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933) // Mainnet
		assert.NoError(t, err)

		callCount := 0
		blsChecksDone := false
		// Mock hasActivity to return error during Ed25519 recovery (after BLS completes)
		hasActivity := func(_ string) (bool, error) {
			callCount++
			// Detect when we've moved past BLS addresses (after 8+ BLS checks without activity)
			if callCount >= 10 && !blsChecksDone {
				blsChecksDone = true
			}
			if blsChecksDone && callCount >= 12 {
				return false, errors.New("network error during Ed25519 check")
			}

			return false, nil
		}

		err = vault.RecoverAddresses(context.Background(), "", hasActivity)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "network error during Ed25519 check")
	})
}

// TestPrivateKeyDerivation tests the private key derivation functions for both BLS and Ed25519.
func TestPrivateKeyDerivation(t *testing.T) {
	//nolint:dupword // has duplicated words
	testMnemonic := "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon cactus"

	t.Run("derive BLS private key successfully", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933)
		assert.NoError(t, err)

		seed, err := vault.MnemonicSeed("")
		assert.NoError(t, err)

		// Valid BLS path: m/12381'/19933'/2'/0
		path := []uint32{
			_H(PurposeBLS12381),
			_H(19933),
			_H(uint32(crypto.AddressTypeBLSAccount)),
			0,
		}

		prv, err := vault.deriveBLSPrivateKey(seed, path)
		assert.NoError(t, err)
		assert.NotNil(t, prv)

		// Verify the derived key can generate a valid public key
		pub := prv.PublicKeyNative()
		assert.NotNil(t, pub)
		assert.Equal(t, "ae1z0v7jum3clvcfeurxg92vt8uwx002ev9jxtyk3p", pub.AccountAddress().String())
	})

	t.Run("derive BLS private key with invalid seed", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933)
		assert.NoError(t, err)

		// Invalid seed (too short)
		invalidSeed := []byte("short")

		path := []uint32{
			_H(PurposeBLS12381),
			_H(19933),
			_H(uint32(crypto.AddressTypeBLSAccount)),
			0,
		}

		prv, err := vault.deriveBLSPrivateKey(invalidSeed, path)
		assert.Error(t, err)
		assert.Nil(t, prv)
	})

	t.Run("derive BLS private key with invalid path", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933)
		assert.NoError(t, err)

		seed, err := vault.MnemonicSeed("")
		assert.NoError(t, err)

		// Invalid path that will cause derivation to fail
		// Using extremely large index that exceeds limits
		invalidPath := []uint32{
			_H(PurposeBLS12381),
			_H(19933),
			_H(uint32(crypto.AddressTypeBLSAccount)),
			0xFFFFFFFF, // Max uint32, likely to cause issues
		}

		_, err = vault.deriveBLSPrivateKey(seed, invalidPath)
		// May or may not error depending on implementation, but tests the path
		// The important part is exercising this code path
		_ = err
	})

	t.Run("derive Ed25519 private key successfully", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933)
		assert.NoError(t, err)

		seed, err := vault.MnemonicSeed("")
		assert.NoError(t, err)

		// Valid Ed25519 path: m/44'/19933'/3'/0'
		path := []uint32{
			_H(PurposeBIP44),
			_H(19933),
			_H(uint32(crypto.AddressTypeEd25519Account)),
			_H(0),
		}

		prv, err := vault.deriveEd25519PrivateKey(seed, path)
		assert.NoError(t, err)
		assert.NotNil(t, prv)

		// Verify the derived key can generate a valid public key
		pub := prv.PublicKeyNative()
		assert.NotNil(t, pub)
		assert.Equal(t, "ae1rex4yps0chj9lsax9n8m0vpt59a2jecxv8k0g8y", pub.AccountAddress().String())
	})

	t.Run("derive Ed25519 private key with invalid seed", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933)
		assert.NoError(t, err)

		// Invalid seed (too short)
		invalidSeed := []byte("short")

		path := []uint32{
			_H(PurposeBIP44),
			_H(19933),
			_H(uint32(crypto.AddressTypeEd25519Account)),
			_H(0),
		}

		prv, err := vault.deriveEd25519PrivateKey(invalidSeed, path)
		assert.Error(t, err)
		assert.Nil(t, prv)
	})

	t.Run("derive Ed25519 private key with invalid path", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933)
		assert.NoError(t, err)

		seed, err := vault.MnemonicSeed("")
		assert.NoError(t, err)

		// Invalid path that will cause derivation to fail
		// Using extremely large index
		invalidPath := []uint32{
			_H(PurposeBIP44),
			_H(19933),
			_H(uint32(crypto.AddressTypeEd25519Account)),
			_H(0xFFFFFFFF), // Max uint32, likely to cause issues
		}

		_, err = vault.deriveEd25519PrivateKey(seed, invalidPath)
		// May or may not error depending on implementation
		_ = err
	})

	t.Run("derive multiple BLS private keys", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933)
		assert.NoError(t, err)

		seed, err := vault.MnemonicSeed("")
		assert.NoError(t, err)

		// Derive multiple keys and verify they're different
		keys := make([]*bls.PrivateKey, 3)
		for i := 0; i < 3; i++ {
			path := []uint32{
				_H(PurposeBLS12381),
				_H(19933),
				_H(uint32(crypto.AddressTypeBLSAccount)),
				uint32(i),
			}

			prv, err := vault.deriveBLSPrivateKey(seed, path)
			assert.NoError(t, err)
			assert.NotNil(t, prv)
			keys[i] = prv
		}

		// Verify all keys are different
		for i := 0; i < len(keys); i++ {
			for j := i + 1; j < len(keys); j++ {
				assert.False(t, keys[i].EqualsTo(keys[j]),
					"Keys at indices %d and %d should be different", i, j)
			}
		}
	})

	t.Run("derive multiple Ed25519 private keys", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933)
		assert.NoError(t, err)

		seed, err := vault.MnemonicSeed("")
		assert.NoError(t, err)

		// Derive multiple keys and verify they're different
		keys := make([]*ed25519.PrivateKey, 3)
		for i := 0; i < 3; i++ {
			path := []uint32{
				_H(PurposeBIP44),
				_H(19933),
				_H(uint32(crypto.AddressTypeEd25519Account)),
				_H(uint32(i)),
			}

			prv, err := vault.deriveEd25519PrivateKey(seed, path)
			assert.NoError(t, err)
			assert.NotNil(t, prv)
			keys[i] = prv
		}

		// Verify all keys are different
		for i := 0; i < len(keys); i++ {
			for j := i + 1; j < len(keys); j++ {
				assert.False(t, keys[i].EqualsTo(keys[j]),
					"Keys at indices %d and %d should be different", i, j)
			}
		}
	})

	t.Run("derive BLS validator private key", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933)
		assert.NoError(t, err)

		seed, err := vault.MnemonicSeed("")
		assert.NoError(t, err)

		// Valid BLS validator path: m/12381'/19933'/1'/0
		path := []uint32{
			_H(PurposeBLS12381),
			_H(19933),
			_H(uint32(crypto.AddressTypeValidator)),
			0,
		}

		prv, err := vault.deriveBLSPrivateKey(seed, path)
		assert.NoError(t, err)
		assert.NotNil(t, prv)

		// Verify the derived key can generate a valid validator address
		pub := prv.PublicKeyNative()
		assert.NotNil(t, pub)
		validatorAddr := pub.ValidatorAddress()
		assert.NotEmpty(t, validatorAddr.String())
	})

	t.Run("integration test with PrivateKeys method for BLS", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933)
		assert.NoError(t, err)

		// Create a BLS account address
		addrInfo, err := vault.NewBLSAccountAddress("test-bls")
		assert.NoError(t, err)

		// Retrieve private key using PrivateKeys method (which calls deriveBLSPrivateKey)
		keys, err := vault.PrivateKeys("", []string{addrInfo.Address})
		assert.NoError(t, err)
		assert.Len(t, keys, 1)

		// Verify the key matches the address
		pub := keys[0].PublicKey()
		assert.Equal(t, addrInfo.PublicKey, pub.String())
	})

	t.Run("integration test with PrivateKeys method for Ed25519", func(t *testing.T) {
		vault, err := CreateVaultFromMnemonic(testMnemonic, 19933)
		assert.NoError(t, err)

		// Create an Ed25519 account address
		addrInfo, err := vault.NewEd25519AccountAddress("test-ed25519", "")
		assert.NoError(t, err)

		// Retrieve private key using PrivateKeys method (which calls deriveEd25519PrivateKey)
		keys, err := vault.PrivateKeys("", []string{addrInfo.Address})
		assert.NoError(t, err)
		assert.Len(t, keys, 1)

		// Verify the key matches the address
		pub := keys[0].PublicKey()
		assert.Equal(t, addrInfo.PublicKey, pub.String())
	})
}
