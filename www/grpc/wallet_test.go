package grpc

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/aerium-network/aerium/crypto"
	"github.com/aerium-network/aerium/types/tx"
	"github.com/aerium-network/aerium/wallet"
	aerium "github.com/aerium-network/aerium/www/grpc/gen/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDisableWallet(t *testing.T) {
	conf := testConfig()
	conf.EnableWallet = false

	td := setup(t, conf)
	conn, client := td.walletClient(t)

	res, err := client.CreateWallet(context.Background(),
		&aerium.CreateWalletRequest{
			WalletName: "TestWallet",
		})
	assert.ErrorIs(t, err, status.Error(codes.Unimplemented, "unknown service aerium.Wallet"))
	assert.Nil(t, res)

	assert.Nil(t, conn.Close(), "Error closing connection")
	td.StopServer()
}

func TestRestoreWallet(t *testing.T) {
	config := testConfig()
	config.EnableWallet = true

	td := setup(t, config)
	conn, client := td.walletClient(t)

	t.Run("should return error if no wallet name provided", func(t *testing.T) {
		res, err := client.RestoreWallet(context.Background(),
			&aerium.RestoreWalletRequest{})
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("should return error if no mnemonic provided", func(t *testing.T) {
		res, err := client.RestoreWallet(context.Background(),
			&aerium.RestoreWalletRequest{
				WalletName: "test",
			})
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("should restore wallet", func(t *testing.T) {
		mnemonic, err := wallet.GenerateMnemonic(128)
		assert.NoError(t, err)

		res, err := client.RestoreWallet(context.Background(),
			&aerium.RestoreWalletRequest{
				WalletName: "test",
				Mnemonic:   mnemonic,
			})
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	assert.Nil(t, conn.Close(), "Error closing connection")
	td.StopServer()
}

func TestCreateWallet(t *testing.T) {
	conf := testConfig()
	conf.EnableWallet = true

	td := setup(t, conf)
	conn, client := td.walletClient(t)

	t.Run("No name, should return an error", func(t *testing.T) {
		res, err := client.CreateWallet(context.Background(),
			&aerium.CreateWalletRequest{
				WalletName: "",
			})
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("Bad name, should return an error", func(t *testing.T) {
		res, err := client.CreateWallet(context.Background(),
			&aerium.CreateWalletRequest{
				WalletName: "..",
			})
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("Should create wallet", func(t *testing.T) {
		res, err := client.CreateWallet(context.Background(),
			&aerium.CreateWalletRequest{
				WalletName: "test",
			})
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	assert.Nil(t, conn.Close(), "Error closing connection")
	td.StopServer()
}

func TestLoadWallet(t *testing.T) {
	conf := testConfig()
	conf.EnableWallet = true

	td := setup(t, conf)
	conn, client := td.walletClient(t)

	wltName := "default_wallet"
	wltAddrInfo, err := td.defaultWallet.NewBLSAccountAddress("test")
	assert.NoError(t, err)
	require.NoError(t, td.defaultWallet.Save())

	t.Run("Load non-existing wallet", func(t *testing.T) {
		res, err := client.LoadWallet(context.Background(),
			&aerium.LoadWalletRequest{
				WalletName: "non-existing",
			})
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("Load existing wallet", func(t *testing.T) {
		res, err := client.LoadWallet(context.Background(),
			&aerium.LoadWalletRequest{
				WalletName: wltName,
			})
		require.NoError(t, err)
		assert.Equal(t, wltName, res.WalletName)
	})

	t.Run("Load wallet again", func(t *testing.T) {
		res, err := client.LoadWallet(context.Background(),
			&aerium.LoadWalletRequest{
				WalletName: wltName,
			})
		require.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("Unload unknown wallet", func(t *testing.T) {
		res, err := client.UnloadWallet(context.Background(),
			&aerium.UnloadWalletRequest{
				WalletName: "not-loade",
			})
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("Sign raw transaction, OK", func(t *testing.T) {
		wltAddr, _ := crypto.AddressFromString(wltAddrInfo.Address)
		bondTx := tx.NewBondTx(td.RandHeight(), wltAddr, td.RandValAddress(), nil, td.RandAmount(),
			td.RandAmount())

		data, err := bondTx.Bytes()
		assert.NoError(t, err)

		res, err := client.SignRawTransaction(context.Background(),
			&aerium.SignRawTransactionRequest{
				WalletName:     wltName,
				RawTransaction: hex.EncodeToString(data),
				Password:       "",
			})
		assert.NoError(t, err)
		assert.Equal(t, bondTx.ID().String(), res.TransactionId)

		signedTx, err := tx.FromBytes(td.DecodingHex(res.SignedRawTransaction))
		assert.NoError(t, err)
		assert.NotNil(t, signedTx.Signature())
		assert.NoError(t, signedTx.BasicCheck())
	})

	t.Run("Sign raw transaction using not loaded wallet", func(t *testing.T) {
		wltAddr, _ := crypto.AddressFromString(wltAddrInfo.Address)
		bondTx := tx.NewBondTx(td.RandHeight(), wltAddr, td.RandValAddress(), nil, td.RandAmount(),
			td.RandAmount())

		data, err := bondTx.Bytes()
		assert.NoError(t, err)

		res, err := client.SignRawTransaction(context.Background(),
			&aerium.SignRawTransactionRequest{
				WalletName:     "not-loaded-wallet",
				RawTransaction: hex.EncodeToString(data),
				Password:       "",
			})
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("Sign invalid raw transaction", func(t *testing.T) {
		res, err := client.SignRawTransaction(context.Background(),
			&aerium.SignRawTransactionRequest{
				WalletName:     wltName,
				RawTransaction: "bad0",
				Password:       "",
			})
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("Unload wallet", func(t *testing.T) {
		res, err := client.UnloadWallet(context.Background(),
			&aerium.UnloadWalletRequest{
				WalletName: wltName,
			})
		require.NoError(t, err)
		assert.Equal(t, wltName, res.WalletName)
	})

	assert.Nil(t, conn.Close(), "Error closing connection")
	td.StopServer()
}

func TestGetValidatorAddress(t *testing.T) {
	conf := testConfig()
	conf.EnableWallet = true

	td := setup(t, conf)
	conn, client := td.walletClient(t)

	t.Run("Invalid public key", func(t *testing.T) {
		res, err := client.GetValidatorAddress(context.Background(),
			&aerium.GetValidatorAddressRequest{PublicKey: "something"})
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("OK", func(t *testing.T) {
		valKey := td.RandValKey()
		pubKey := valKey.PublicKey()

		res, err := client.GetValidatorAddress(context.Background(),
			&aerium.GetValidatorAddressRequest{PublicKey: pubKey.String()})

		assert.Nil(t, err)
		assert.Equal(t, pubKey.ValidatorAddress().String(), res.Address)
	})

	assert.Nil(t, conn.Close(), "Error closing connection")
	td.StopServer()
}

func TestGetTotalBalance(t *testing.T) {
	conf := testConfig()
	conf.EnableWallet = true

	td := setup(t, conf)
	conn, client := td.walletClient(t)

	t.Run("wallet not loaded", func(t *testing.T) {
		res, err := client.GetTotalBalance(context.Background(),
			&aerium.GetTotalBalanceRequest{})
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("OK", func(t *testing.T) {
		walletName := "default_wallet"
		_, err := client.LoadWallet(context.Background(), &aerium.LoadWalletRequest{
			WalletName: walletName,
		})
		assert.NoError(t, err)

		res, err := client.GetTotalBalance(context.Background(),
			&aerium.GetTotalBalanceRequest{
				WalletName: walletName,
			})
		assert.NoError(t, err)
		assert.Equal(t, walletName, res.WalletName)
		assert.Zero(t, res.TotalBalance)
	})

	assert.Nil(t, conn.Close(), "Error closing connection")
	td.StopServer()
}

func TestGetWalletInfo_DefaultFee(t *testing.T) {
	conf := testConfig()
	conf.EnableWallet = true

	td := setup(t, conf)
	conn, client := td.walletClient(t)
	defer func() {
		assert.Nil(t, conn.Close(), "Error closing connection")
		td.StopServer()
	}()

	walletName := "default_wallet"

	// load the default wallet first
	_, err := client.LoadWallet(context.Background(), &aerium.LoadWalletRequest{
		WalletName: walletName,
	})
	require.NoError(t, err)

	// fetch wallet info and verify default_fee is exposed
	res, err := client.GetWalletInfo(context.Background(), &aerium.GetWalletInfoRequest{
		WalletName: walletName,
	})
	require.NoError(t, err)
	require.NotNil(t, res)
	// Default fee should be the vault default (0.01 AUM => 10_000_000 NanoAUM)
	assert.EqualValues(t, int64(10_000_000), res.DefaultFee)
}

func TestGetNewAddress(t *testing.T) {
	conf := testConfig()
	conf.EnableWallet = true

	td := setup(t, conf)
	conn, client := td.walletClient(t)

	wltName := "default_wallet"
	_, err := td.defaultWallet.NewBLSAccountAddress("test")
	assert.NoError(t, err)
	require.NoError(t, td.defaultWallet.Save())

	t.Run("New address with BLS account", func(t *testing.T) {
		_, err = client.LoadWallet(context.Background(),
			&aerium.LoadWalletRequest{
				WalletName: wltName,
			})
		require.NoError(t, err)

		res, err := client.GetNewAddress(context.Background(),
			&aerium.GetNewAddressRequest{
				WalletName:  wltName,
				AddressType: aerium.AddressType_ADDRESS_TYPE_BLS_ACCOUNT,
				Label:       "bls-account",
			})
		assert.Nil(t, err)
		assert.Equal(t, wltName, res.WalletName)
		assert.NotEmpty(t, res.AddressInfo.PublicKey)
		assert.NotEmpty(t, res.AddressInfo.Path)
		assert.Equal(t, "bls-account", res.AddressInfo.Label)

		_, err = client.UnloadWallet(context.Background(),
			&aerium.UnloadWalletRequest{
				WalletName: wltName,
			})
		require.NoError(t, err)
	})

	t.Run("New address with validator account", func(t *testing.T) {
		_, err = client.LoadWallet(context.Background(),
			&aerium.LoadWalletRequest{
				WalletName: wltName,
			})
		require.NoError(t, err)

		res, err := client.GetNewAddress(context.Background(),
			&aerium.GetNewAddressRequest{
				WalletName:  wltName,
				AddressType: aerium.AddressType_ADDRESS_TYPE_VALIDATOR,
				Label:       "validator",
			})
		assert.Nil(t, err)
		assert.Equal(t, wltName, res.WalletName)
		assert.NotEmpty(t, res.AddressInfo.PublicKey)
		assert.NotEmpty(t, res.AddressInfo.Path)
		assert.Equal(t, "validator", res.AddressInfo.Label)

		_, err = client.UnloadWallet(context.Background(),
			&aerium.UnloadWalletRequest{
				WalletName: wltName,
			})
		require.NoError(t, err)
	})

	t.Run("Error with new address with treasury", func(t *testing.T) {
		_, err = client.LoadWallet(context.Background(),
			&aerium.LoadWalletRequest{
				WalletName: wltName,
			})
		require.NoError(t, err)

		res, err := client.GetNewAddress(context.Background(),
			&aerium.GetNewAddressRequest{
				WalletName:  wltName,
				AddressType: aerium.AddressType_ADDRESS_TYPE_TREASURY,
				Label:       "treasury",
			})
		assert.NotNil(t, err)
		assert.Nil(t, res)

		_, err = client.UnloadWallet(context.Background(),
			&aerium.UnloadWalletRequest{
				WalletName: wltName,
			})
		require.NoError(t, err)
	})

	assert.Nil(t, conn.Close(), "Error closing connection")
	td.StopServer()
}
