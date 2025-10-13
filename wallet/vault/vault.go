package vault

import (
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"maps"

	"github.com/aerium-network/aerium/crypto"
	"github.com/aerium-network/aerium/crypto/bls"
	blshdkeychain "github.com/aerium-network/aerium/crypto/bls/hdkeychain"
	"github.com/aerium-network/aerium/crypto/ed25519"
	ed25519hdkeychain "github.com/aerium-network/aerium/crypto/ed25519/hdkeychain"
	"github.com/aerium-network/aerium/types/amount"
	"github.com/aerium-network/aerium/util/bip39"
	"github.com/aerium-network/aerium/wallet/addresspath"
	"github.com/aerium-network/aerium/wallet/encrypter"
	"golang.org/x/exp/slices"
)

//
// Deterministic Hierarchical Derivation Path
//
// Overview:
//
// This specification defines a hierarchical derivation path for generating addresses, based on BIP32.
// The path is structured into four distinct levels:
//
// m / purpose' / coin_type' / address_type' / address_index
//
// Explanation:
//
//   `m` Denotes the master node (or root) of the tree
//   `'` Apostrophe in the path indicates that BIP32 hardened derivation is used.
//   `/` Separates the tree into depths, thus i / j signifies that j is a child of i
//
// Path Components:
//
// * `purpose`: Indicates the specific use case for the derived addresses:
//    - 12381: Used for the BLS12-381 curve.
//    - 65535: Used for imported private keys.
//    - 44: A comprehensive purpose for standard curves, based on BIP-44.
//
// * `coin_type`: Identifies the coin type:
//    - 19933: Aerium Mainnet
//    - 19944: Aerium Testnet
//
// * `address_type`: Specifies the type of address.
//
// * `address_index`: A sequential number and increase when a new address is derived.
//
// References:
//  - https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki
//

// VaultType represents the type of vault.
type VaultType int

const (
	TypeFull     VaultType = iota + 1 // Full vault with private keys.
	TypeNeutered                      // Neutered vault without private keys.
)

// String returns the string representation of the VaultType.
func (vt VaultType) String() string {
	switch vt {
	case TypeFull:
		return "Full"
	case TypeNeutered:
		return "Neutered"
	default:
		return "Unknown"
	}
}

type AddressInfo struct {
	Address   string `json:"address"`    // Address in the wallet
	PublicKey string `json:"public_key"` // Public key associated with the address
	Label     string `json:"label"`      // Label for the address
	Path      string `json:"path"`       // Path for the address
}

const (
	PurposeBLS12381         = uint32(12381)
	PurposeBIP44            = uint32(44)
	PurposeImportPrivateKey = uint32(65535)

	PurposeBLS12381Hardened         = PurposeBLS12381 + addresspath.HardenedKeyStart
	PurposeBIP44Hardened            = PurposeBIP44 + addresspath.HardenedKeyStart
	PurposeImportPrivateKeyHardened = PurposeImportPrivateKey + addresspath.HardenedKeyStart
)

// AddressGapLimit is the maximum number of consecutive inactive addresses before stopping recovery.
const AddressGapLimit = 8

type Vault struct {
	Type       VaultType              `json:"type"`        // Vault type: Full or Neutered
	CoinType   uint32                 `json:"coin_type"`   // Coin type: 21888 for Mainnet, 21777 for Testnet
	DefaultFee amount.Amount          `json:"default_fee"` // The Vault's default fee
	Addresses  map[string]AddressInfo `json:"addresses"`   // All addresses that are stored in the vault
	Encrypter  encrypter.Encrypter    `json:"encrypter"`   // Encryption algorithm
	KeyStore   string                 `json:"key_store"`   // KeyStore that stores the secrets and encrypts using Encrypter
	Purposes   purposes               `json:"purposes"`    // Contains Purposes of the vault
}

type keyStore struct {
	MasterNode   masterNode `json:"master_node"`   // HD Root Tree (Master node)
	ImportedKeys []string   `json:"imported_keys"` // Imported private keys
}

type masterNode struct {
	Mnemonic string `json:"seed,omitempty"` // Seed phrase or mnemonic (encrypted)
}

type purposes struct {
	PurposeBLS   purposeBLS   `json:"purpose_bls"`   // BLS Purpose: m/12381'/21888'/1' or 2'/0
	PurposeBIP44 purposeBIP44 `json:"purpose_bip44"` // BIP44 Purpose: m/44'/21888'/3'/0'
}

type purposeBLS struct {
	XPubValidator      string `json:"xpub_account"`         // Extended public key for account: m/12381'/21888'/1'/0
	XPubAccount        string `json:"xpub_validator"`       // Extended public key for validator: m/12381'/21888'/2'/0
	NextAccountIndex   uint32 `json:"next_account_index"`   // Index of next derived account
	NextValidatorIndex uint32 `json:"next_validator_index"` // Index of next derived validator
}

type purposeBIP44 struct {
	NextEd25519Index uint32 `json:"next_ed25519_index"` // Index of next Ed25519 derived account: m/44'/21888/3'/0'
}

func CreateVaultFromMnemonic(mnemonic string, coinType uint32) (*Vault, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil, err
	}
	masterKey, err := blshdkeychain.NewMaster(seed, false)
	if err != nil {
		return nil, err
	}
	enc := encrypter.NopeEncrypter()

	xPubValidator, err := masterKey.DerivePath([]uint32{
		_H(PurposeBLS12381),
		_H(coinType),
		_H(crypto.AddressTypeValidator),
	})
	if err != nil {
		return nil, err
	}

	xPubAccount, err := masterKey.DerivePath([]uint32{
		_H(PurposeBLS12381),
		_H(coinType),
		_H(crypto.AddressTypeBLSAccount),
	})
	if err != nil {
		return nil, err
	}

	store := keyStore{
		MasterNode: masterNode{
			Mnemonic: mnemonic,
		},
		ImportedKeys: make([]string, 0),
	}

	storeDate, err := json.Marshal(store)
	if err != nil {
		return nil, err
	}

	return &Vault{
		Type:       TypeFull,
		CoinType:   coinType,
		Encrypter:  enc,
		Addresses:  make(map[string]AddressInfo),
		KeyStore:   string(storeDate),
		DefaultFee: amount.Amount(10_000_000), // 0.01 AUM
		Purposes: purposes{
			PurposeBLS: purposeBLS{
				XPubValidator: xPubValidator.Neuter().String(),
				XPubAccount:   xPubAccount.Neuter().String(),
			},
		},
	}, nil
}

func (v *Vault) Neuter() *Vault {
	neutered := &Vault{
		Type:      TypeNeutered,
		CoinType:  v.CoinType,
		Encrypter: encrypter.NopeEncrypter(),
		Addresses: make(map[string]AddressInfo),
		KeyStore:  "",
		Purposes:  v.Purposes,
	}

	maps.Copy(neutered.Addresses, v.Addresses)

	return neutered
}

func (v *Vault) IsNeutered() bool {
	return v.Type == TypeNeutered
}

func (v *Vault) UpdatePassword(oldPassword, newPassword string, opts ...encrypter.Option) error {
	if v.IsNeutered() {
		return ErrNeutered
	}

	keyStore, err := v.decryptKeyStore(oldPassword)
	if err != nil {
		return err
	}

	newEncrypter := encrypter.NopeEncrypter()
	if newPassword != "" {
		newEncrypter = encrypter.DefaultEncrypter(opts...)
	}
	v.Encrypter = newEncrypter
	err = v.encryptKeyStore(keyStore, newPassword)
	if err != nil {
		return err
	}

	v.Encrypter = newEncrypter

	return nil
}

func (v *Vault) Label(addr string) string {
	info, ok := v.Addresses[addr]
	if !ok {
		return ""
	}

	return info.Label
}

func (v *Vault) SetLabel(addr, label string) error {
	info, ok := v.Addresses[addr]
	if !ok {
		return NewErrAddressNotFound(addr)
	}

	info.Label = label
	v.Addresses[addr] = info

	return nil
}

func (v *Vault) AddressInfos() []AddressInfo {
	addrs := make([]AddressInfo, 0, 1)
	for _, addrInfo := range v.Addresses {
		addrs = append(addrs, addrInfo)
	}

	v.sortAddressesByAddressIndex(addrs...)
	v.sortAddressesByAddressType(addrs...)
	v.sortAddressesByPurpose(addrs...)

	return addrs
}

func (v *Vault) AllValidatorAddresses() []AddressInfo {
	addrs := make([]AddressInfo, 0, v.AddressCount()/2)
	for _, addrInfo := range v.Addresses {
		addrPath, _ := addresspath.FromString(addrInfo.Path)
		if addrPath.AddressType() == _H(crypto.AddressTypeValidator) {
			addrs = append(addrs, addrInfo)
		}
	}

	v.sortAddressesByAddressIndex(addrs...)
	v.sortAddressesByPurpose(addrs...)

	return addrs
}

func (v *Vault) AllAccountAddresses() []AddressInfo {
	addrs := make([]AddressInfo, 0, v.AddressCount()/2)
	for _, addrInfo := range v.Addresses {
		addrPath, _ := addresspath.FromString(addrInfo.Path)
		if addrPath.AddressType() != _H(crypto.AddressTypeValidator) {
			addrs = append(addrs, addrInfo)
		}
	}

	v.sortAddressesByAddressIndex(addrs...)
	v.sortAddressesByPurpose(addrs...)

	return addrs
}

func (*Vault) sortAddressesByPurpose(addrs ...AddressInfo) {
	slices.SortStableFunc(addrs, func(a, b AddressInfo) int {
		pathA, _ := addresspath.FromString(a.Path)
		pathB, _ := addresspath.FromString(b.Path)

		return cmp.Compare(pathA.Purpose(), pathB.Purpose())
	})
}

func (*Vault) sortAddressesByAddressType(addrs ...AddressInfo) {
	slices.SortStableFunc(addrs, func(a, b AddressInfo) int {
		pathA, _ := addresspath.FromString(a.Path)
		pathB, _ := addresspath.FromString(b.Path)

		return cmp.Compare(pathA.AddressType(), pathB.AddressType())
	})
}

func (*Vault) sortAddressesByAddressIndex(addrs ...AddressInfo) {
	slices.SortStableFunc(addrs, func(a, b AddressInfo) int {
		pathA, _ := addresspath.FromString(a.Path)
		pathB, _ := addresspath.FromString(b.Path)

		return cmp.Compare(pathA.AddressIndex(), pathB.AddressIndex())
	})
}

func (v *Vault) IsEncrypted() bool {
	return v.Encrypter.IsEncrypted()
}

func (v *Vault) AddressCount() int {
	return len(v.Addresses)
}

func (v *Vault) AddressFromPath(p string) *AddressInfo {
	for _, addressInfo := range v.Addresses {
		if addressInfo.Path == p {
			return &addressInfo
		}
	}

	return nil
}

func (v *Vault) ImportBLSPrivateKey(password string, prv *bls.PrivateKey) error {
	if v.IsNeutered() {
		return ErrNeutered
	}

	keyStore, err := v.decryptKeyStore(password)
	if err != nil {
		return err
	}

	addressIndex := len(keyStore.ImportedKeys)
	pub := prv.PublicKeyNative()

	accAddr := pub.AccountAddress()
	if v.Contains(accAddr.String()) {
		return ErrAddressExists
	}

	valAddr := pub.ValidatorAddress()
	if v.Contains(valAddr.String()) {
		return ErrAddressExists
	}

	blsAccPathStr := addresspath.NewPath(
		_H(PurposeImportPrivateKey),
		_H(v.CoinType),
		_H(crypto.AddressTypeBLSAccount),
		_H(addressIndex)).String()

	blsValidatorPathStr := addresspath.NewPath(
		_H(PurposeImportPrivateKey),
		_H(v.CoinType),
		_H(crypto.AddressTypeValidator),
		_H(addressIndex)).String()

	v.Addresses[accAddr.String()] = AddressInfo{
		Address:   accAddr.String(),
		PublicKey: pub.String(),
		Label:     "Imported BLS Account Address",
		Path:      blsAccPathStr,
	}

	v.Addresses[valAddr.String()] = AddressInfo{
		Address:   valAddr.String(),
		PublicKey: pub.String(),
		Label:     "Imported Validator Address",
		Path:      blsValidatorPathStr,
	}

	keyStore.ImportedKeys = append(keyStore.ImportedKeys, prv.String())

	err = v.encryptKeyStore(keyStore, password)
	if err != nil {
		return err
	}

	return nil
}

func (v *Vault) ImportEd25519PrivateKey(password string, prv *ed25519.PrivateKey) error {
	if v.IsNeutered() {
		return ErrNeutered
	}

	keyStore, err := v.decryptKeyStore(password)
	if err != nil {
		return err
	}

	addressIndex := len(keyStore.ImportedKeys)
	pub := prv.PublicKeyNative()

	accAddr := pub.AccountAddress()
	if v.Contains(accAddr.String()) {
		return ErrAddressExists
	}

	accPathStr := addresspath.NewPath(
		_H(PurposeImportPrivateKey),
		_H(v.CoinType),
		_H(crypto.AddressTypeEd25519Account),
		_H(addressIndex)).String()

	v.Addresses[accAddr.String()] = AddressInfo{
		Address:   accAddr.String(),
		PublicKey: pub.String(),
		Label:     "Imported Ed25519 Account Address",
		Path:      accPathStr,
	}

	keyStore.ImportedKeys = append(keyStore.ImportedKeys, prv.String())

	err = v.encryptKeyStore(keyStore, password)
	if err != nil {
		return err
	}

	return nil
}

// PrivateKeys retrieves the private keys for the given addresses using the provided password.
func (v *Vault) PrivateKeys(password string, addrs []string) ([]crypto.PrivateKey, error) {
	if v.IsNeutered() {
		return nil, ErrNeutered
	}

	// Decrypt the key store once to avoid decrypting for each key.
	keyStore, err := v.decryptKeyStore(password)
	if err != nil {
		return nil, err
	}
	seed := bip39.NewSeed(keyStore.MasterNode.Mnemonic, "")

	keys := make([]crypto.PrivateKey, len(addrs))
	for i, addr := range addrs {
		info := v.AddressInfo(addr)
		if info == nil {
			return nil, NewErrAddressNotFound(addr)
		}

		hdPath, err := addresspath.FromString(info.Path)
		if err != nil {
			return nil, err
		}

		if hdPath.CoinType() != _H(v.CoinType) {
			return nil, ErrInvalidCoinType
		}

		switch hdPath.Purpose() {
		case _H(PurposeBLS12381):
			prvKey, err := v.deriveBLSPrivateKey(seed, hdPath)
			if err != nil {
				return nil, err
			}
			keys[i] = prvKey
		case _H(PurposeBIP44):
			prvKey, err := v.deriveEd25519PrivateKey(seed, hdPath)
			if err != nil {
				return nil, err
			}
			keys[i] = prvKey
		case _H(PurposeImportPrivateKey):
			index := _N(hdPath.AddressIndex())
			str := keyStore.ImportedKeys[index]

			var prv crypto.PrivateKey
			switch _N(hdPath.AddressType()) {
			case uint32(crypto.AddressTypeValidator),
				uint32(crypto.AddressTypeBLSAccount):
				prv, err = bls.PrivateKeyFromString(str)
				if err != nil {
					return nil, err
				}

			case uint32(crypto.AddressTypeEd25519Account):
				prv, err = ed25519.PrivateKeyFromString(str)
				if err != nil {
					return nil, err
				}
			}

			keys[i] = prv
		default:
			return nil, ErrUnsupportedPurpose
		}
	}

	return keys, nil
}

func (v *Vault) NewValidatorAddress(label string) (*AddressInfo, error) {
	ext, err := blshdkeychain.NewKeyFromString(v.Purposes.PurposeBLS.XPubValidator)
	if err != nil {
		return nil, err
	}
	index := v.Purposes.PurposeBLS.NextValidatorIndex
	ext, err = ext.DerivePath([]uint32{index})
	if err != nil {
		return nil, err
	}

	blsPubKey, err := bls.PublicKeyFromBytes(ext.RawPublicKey())
	if err != nil {
		return nil, err
	}

	addr := blsPubKey.ValidatorAddress().String()
	info := AddressInfo{
		Address:   addr,
		Label:     label,
		PublicKey: blsPubKey.String(),
		Path:      addresspath.NewPath(ext.Path()...).String(),
	}
	v.Addresses[addr] = info
	v.Purposes.PurposeBLS.NextValidatorIndex++

	return &info, nil
}

func (v *Vault) NewBLSAccountAddress(label string) (*AddressInfo, error) {
	ext, err := blshdkeychain.NewKeyFromString(v.Purposes.PurposeBLS.XPubAccount)
	if err != nil {
		return nil, err
	}
	index := v.Purposes.PurposeBLS.NextAccountIndex
	info, err := v.deriveBLSAccountAddressAt(ext, index, label)
	if err != nil {
		return nil, err
	}

	v.Addresses[info.Address] = *info
	v.Purposes.PurposeBLS.NextAccountIndex++

	return info, nil
}

func (*Vault) deriveBLSAccountAddressAt(ext *blshdkeychain.ExtendedKey,
	index uint32, label string,
) (*AddressInfo, error) {
	ext, err := ext.DerivePath([]uint32{index})
	if err != nil {
		return nil, err
	}

	blsPubKey, err := bls.PublicKeyFromBytes(ext.RawPublicKey())
	if err != nil {
		return nil, err
	}

	addr := blsPubKey.AccountAddress().String()
	info := AddressInfo{
		Address:   addr,
		Label:     label,
		PublicKey: blsPubKey.String(),
		Path:      addresspath.NewPath(ext.Path()...).String(),
	}

	return &info, nil
}

func (v *Vault) NewEd25519AccountAddress(label, password string) (*AddressInfo, error) {
	seed, err := v.MnemonicSeed(password)
	if err != nil {
		return nil, err
	}

	masterKey, err := ed25519hdkeychain.NewMaster(seed)
	if err != nil {
		return nil, err
	}

	index := v.Purposes.PurposeBIP44.NextEd25519Index
	info, err := v.deriveEd25519AccountAddressAt(masterKey, index, label)
	if err != nil {
		return nil, err
	}
	v.Addresses[info.Address] = *info
	v.Purposes.PurposeBIP44.NextEd25519Index++

	return info, nil
}

func (v *Vault) deriveEd25519AccountAddressAt(masterKey *ed25519hdkeychain.ExtendedKey,
	index uint32, label string,
) (*AddressInfo, error) {
	ext, err := masterKey.DerivePath([]uint32{
		_H(PurposeBIP44),
		_H(v.CoinType),
		_H(crypto.AddressTypeEd25519Account),
		_H(index),
	})
	if err != nil {
		return nil, err
	}

	ed25519PubKey, err := ed25519.PublicKeyFromBytes(ext.RawPublicKey())
	if err != nil {
		return nil, err
	}

	addr := ed25519PubKey.AccountAddress().String()
	info := AddressInfo{
		Address:   addr,
		Label:     label,
		PublicKey: ed25519PubKey.String(),
		Path:      addresspath.NewPath(ext.Path()...).String(),
	}

	return &info, nil
}

// AddressInfo like it can return bls.PublicKey instead of string.
func (v *Vault) AddressInfo(addr string) *AddressInfo {
	info, ok := v.Addresses[addr]
	if !ok {
		return nil
	}

	return &info
}

func (v *Vault) Contains(addr string) bool {
	return v.AddressInfo(addr) != nil
}

func (v *Vault) Mnemonic(password string) (string, error) {
	keyStore, err := v.decryptKeyStore(password)
	if err != nil {
		return "", err
	}

	return keyStore.MasterNode.Mnemonic, nil
}

func (v *Vault) MnemonicSeed(password string) ([]byte, error) {
	mnemonic, err := v.Mnemonic(password)
	if err != nil {
		return nil, err
	}
	seed := bip39.NewSeed(mnemonic, "")

	return seed, nil
}

func (v *Vault) decryptKeyStore(password string) (*keyStore, error) {
	if v.IsNeutered() {
		return nil, ErrNeutered
	}

	keyStoreData, err := v.Encrypter.Decrypt(v.KeyStore, password)
	if err != nil {
		return nil, err
	}

	keyStore := new(keyStore)
	err = json.Unmarshal([]byte(keyStoreData), keyStore)
	if err != nil {
		return nil, err
	}

	return keyStore, nil
}

func (v *Vault) encryptKeyStore(keyStore *keyStore, password string) error {
	keyStoreData, err := json.Marshal(keyStore)
	if err != nil {
		return err
	}

	keyStoreEnc, err := v.Encrypter.Encrypt(string(keyStoreData), password)
	if err != nil {
		return err
	}
	v.KeyStore = keyStoreEnc

	return nil
}

func (*Vault) deriveBLSPrivateKey(mnemonicSeed []byte, path []uint32) (*bls.PrivateKey, error) {
	masterKey, err := blshdkeychain.NewMaster(mnemonicSeed, false)
	if err != nil {
		return nil, err
	}
	ext, err := masterKey.DerivePath(path)
	if err != nil {
		return nil, err
	}
	prvBytes, err := ext.RawPrivateKey()
	if err != nil {
		return nil, err
	}

	return bls.PrivateKeyFromBytes(prvBytes)
}

func (*Vault) deriveEd25519PrivateKey(mnemonicSeed []byte, path []uint32) (*ed25519.PrivateKey, error) {
	masterKey, err := ed25519hdkeychain.NewMaster(mnemonicSeed)
	if err != nil {
		return nil, err
	}
	ext, err := masterKey.DerivePath(path)
	if err != nil {
		return nil, err
	}
	prvBytes := ext.RawPrivateKey()

	return ed25519.PrivateKeyFromBytes(prvBytes)
}

// RecoverAddresses restores all previously used addresses when a wallet is recovered from a mnemonic phrase.
// This function attempts to recover both BLS and Ed25519 account addresses. If the wallet is empty,
// Ed25519 is used as the default address type.
//
// An address is considered active if its public key is present in the blockchain database.
// The hasActivity callback should return true if the address has been used previously.
//
// Limitation: If more than 8 consecutive unused addresses exist between used addresses,
// automatic recovery will stop at the gap.
// In such cases, manual address creation is required to recover additional addresses.
func (v *Vault) RecoverAddresses(ctx context.Context, password string,
	hasActivity func(addr string) (bool, error),
) error {
	err := v.recoverBLSAccountAddresses(ctx, hasActivity)
	if err != nil {
		return err
	}

	return v.recoverEd25519AccountAddresses(ctx, password, hasActivity)
}

// recoverBLSAccountAddresses recovers BLS account addresses.
func (v *Vault) recoverBLSAccountAddresses(ctx context.Context, hasActivity func(addr string) (bool, error)) error {
	ext, err := blshdkeychain.NewKeyFromString(v.Purposes.PurposeBLS.XPubAccount)
	if err != nil {
		return err
	}

	lastActiveIndex := -1
	inactiveCount := 0
	for currentIndex := uint32(0); ; currentIndex++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		info, err := v.deriveBLSAccountAddressAt(ext, currentIndex, "")
		if err != nil {
			// This can happen if we try to derive past the key limit.
			// We can break here as we've likely scanned far enough.
			break
		}

		isActive, err := hasActivity(info.Address)
		if err != nil {
			return err
		}

		if isActive {
			lastActiveIndex = int(currentIndex)
			inactiveCount = 0
		} else {
			inactiveCount++
			if inactiveCount >= AddressGapLimit {
				break
			}
		}
	}

	if lastActiveIndex == -1 {
		return nil
	}

	recoveredCount := lastActiveIndex + 1
	// Recover all addresses up to the total number of recovered addresses.
	for i := uint32(0); i < uint32(recoveredCount); i++ {
		if _, err := v.NewBLSAccountAddress(fmt.Sprintf("BLS Account Address %d", i)); err != nil {
			return err
		}
	}

	return nil
}

// recoverEd25519AccountAddresses recovers Ed25519 account addresses.
func (v *Vault) recoverEd25519AccountAddresses(ctx context.Context, password string,
	hasActivity func(addr string) (bool, error),
) error {
	seed, err := v.MnemonicSeed(password)
	if err != nil {
		return err
	}

	masterKey, err := ed25519hdkeychain.NewMaster(seed)
	if err != nil {
		return err
	}

	lastActiveIndex := -1
	inactiveCount := 0
	for currentIndex := uint32(0); ; currentIndex++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		info, err := v.deriveEd25519AccountAddressAt(masterKey, currentIndex, "")
		if err != nil {
			return err
		}

		isActive, err := hasActivity(info.Address)
		if err != nil {
			return err
		}

		if isActive {
			lastActiveIndex = int(currentIndex)
			inactiveCount = 0
		} else {
			inactiveCount++
			if inactiveCount >= AddressGapLimit {
				break
			}
		}
	}

	recoveredCount := 1 // Always recover at least the first address.
	if lastActiveIndex > -1 {
		recoveredCount = lastActiveIndex + 1
	}

	// Recover all addresses up to the total number of recovered addresses.
	for i := uint32(0); i < uint32(recoveredCount); i++ {
		if _, err := v.NewEd25519AccountAddress(fmt.Sprintf("Ed25519 Account Address %d", i), password); err != nil {
			return err
		}
	}

	return nil
}
