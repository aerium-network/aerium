package wallet

import (
	"encoding/json"
	"hash/crc32"
	"time"

	"github.com/aerium-network/aerium/crypto"
	"github.com/aerium-network/aerium/genesis"
	"github.com/aerium-network/aerium/util"
	"github.com/aerium-network/aerium/wallet/vault"
	"github.com/google/uuid"
)

const (
	Version1 = 1 // Initial version

	VersionLatest = Version1
)

type Store struct {
	Version   int               `json:"version"`
	UUID      uuid.UUID         `json:"uuid"`
	CreatedAt time.Time         `json:"created_at"`
	Network   genesis.ChainType `json:"network"`
	VaultCRC  uint32            `json:"crc"`
	Vault     *vault.Vault      `json:"vault"`
	History   history           `json:"history"`
}

func FromBytes(data []byte) (*Store, error) {
	s := new(Store)
	if err := json.Unmarshal(data, s); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Store) ToBytes() ([]byte, error) {
	s.VaultCRC = s.calcVaultCRC()

	return json.MarshalIndent(s, "  ", "  ")
}

func (s *Store) Clone() *Store {
	clonedVault := *s.Vault // Assuming Vault has proper pointer handling internally
	clonedHistory := s.History

	return &Store{
		Version:   s.Version,
		UUID:      s.UUID,
		CreatedAt: s.CreatedAt,
		Network:   s.Network,
		VaultCRC:  s.VaultCRC,
		Vault:     &clonedVault,
		History:   clonedHistory,
	}
}

func (s *Store) ValidateCRC() error {
	crc := s.calcVaultCRC()
	if s.VaultCRC != crc {
		return CRCNotMatchError{
			Expected: crc,
			Got:      s.VaultCRC,
		}
	}

	return nil
}

func (s *Store) UpgradeWallet(walletPath string) error {
	if !s.Network.IsMainnet() {
		crypto.ToTestnetHRP()
	}

	oldVersion := s.Version
	switch oldVersion {
	case Version1:
	default:
		return UnsupportedVersionError{
			WalletVersion:    s.Version,
			SupportedVersion: VersionLatest,
		}
	}

	// Write wallet data.
	s.VaultCRC = s.calcVaultCRC()

	bs, err := s.ToBytes()
	if err != nil {
		return err
	}

	err = util.WriteFile(walletPath, bs)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) calcVaultCRC() uint32 {
	d, err := json.Marshal(s.Vault)
	if err != nil {
		return 0
	}

	return crc32.ChecksumIEEE(d)
}
