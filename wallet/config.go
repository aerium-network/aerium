package wallet

import (
	"github.com/aerium-network/aerium/genesis"
)

type Config struct {
	// private config
	WalletsDir string            `toml:"-"`
	ChainType  genesis.ChainType `toml:"-"`
}

func DefaultConfig() *Config {
	return &Config{}
}
