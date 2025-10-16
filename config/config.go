package config

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aerium-network/aerium/consensus"
	"github.com/aerium-network/aerium/crypto"
	"github.com/aerium-network/aerium/network"
	"github.com/aerium-network/aerium/store"
	"github.com/aerium-network/aerium/sync"
	"github.com/aerium-network/aerium/txpool"
	"github.com/aerium-network/aerium/util"
	"github.com/aerium-network/aerium/util/logger"
	"github.com/aerium-network/aerium/wallet"
	"github.com/aerium-network/aerium/www/grpc"
	"github.com/aerium-network/aerium/www/html"
	"github.com/aerium-network/aerium/www/http"
	"github.com/aerium-network/aerium/www/jsonrpc"
	"github.com/aerium-network/aerium/www/zmq"
	"github.com/pelletier/go-toml/v2"
)

var (
	//go:embed example_config.toml
	exampleConfigBytes []byte

	//go:embed bootstrap.json
	bootstrapInfoBytes []byte

	//go:embed banned_addrs.json
	bannedAddrBytes []byte
)

type Config struct {
	Node    *NodeConfig     `toml:"node"`
	Store   *store.Config   `toml:"store"`
	Network *network.Config `toml:"network"`
	Sync    *sync.Config    `toml:"sync"`
	TxPool  *txpool.Config  `toml:"tx_pool"`
	Logger  *logger.Config  `toml:"logger"`
	GRPC    *grpc.Config    `toml:"grpc"`
	JSONRPC *jsonrpc.Config `toml:"jsonrpc"`
	HTTP    *http.Config    `toml:"http"`
	HTML    *html.Config    `toml:"html"`
	ZeroMq  *zmq.Config     `toml:"zeromq"`

	Consensus     *consensus.Config `toml:"-"`
	WalletManager *wallet.Config    `toml:"-"`
}

type BootstrapInfo struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Website string `json:"website"`
	Address string `json:"address"`
}

type NodeConfig struct {
	RewardAddresses []string `toml:"reward_addresses"`
}

func DefaultNodeConfig() *NodeConfig {
	return &NodeConfig{
		RewardAddresses: []string{},
	}
}

// BasicCheck performs basic checks on the configuration.
func (conf *NodeConfig) BasicCheck() error {
	for _, addrStr := range conf.RewardAddresses {
		addr, err := crypto.AddressFromString(addrStr)
		if err != nil {
			return NodeConfigError{
				Reason: fmt.Sprintf("invalid reward address: %v", err.Error()),
			}
		}

		if !addr.IsAccountAddress() {
			return NodeConfigError{
				Reason: fmt.Sprintf("reward address is not an account address: %s", addrStr),
			}
		}
	}

	return nil
}

func defaultConfig() *Config {
	conf := &Config{
		Node:          DefaultNodeConfig(),
		Store:         store.DefaultConfig(),
		Network:       network.DefaultConfig(),
		Sync:          sync.DefaultConfig(),
		TxPool:        txpool.DefaultConfig(),
		Consensus:     consensus.DefaultConfig(),
		Logger:        logger.DefaultConfig(),
		GRPC:          grpc.DefaultConfig(),
		HTML:          html.DefaultConfig(),
		HTTP:          http.DefaultConfig(),
		JSONRPC:       jsonrpc.DefaultConfig(),
		ZeroMq:        zmq.DefaultConfig(),
		WalletManager: wallet.DefaultConfig(),
	}

	return conf
}

func DefaultConfigMainnet() *Config {
	conf := defaultConfig()

	bootstrapNodes := make([]BootstrapInfo, 0)
	if err := json.Unmarshal(bootstrapInfoBytes, &bootstrapNodes); err != nil {
		panic(err)
	}

	bootstrapAddrs := []string{}
	for _, node := range bootstrapNodes {
		bootstrapAddrs = append(bootstrapAddrs, node.Address)
	}

	bannedList := make([]string, 0)
	if err := json.Unmarshal(bannedAddrBytes, &bannedList); err != nil {
		panic(err)
	}

	bannedAddrs := make(map[crypto.Address]bool)
	for _, str := range bannedList {
		addr, err := crypto.AddressFromString(str)
		if err != nil {
			panic(err)
		}
		bannedAddrs[addr] = true
	}

	conf.Store.BannedAddrs = bannedAddrs
	conf.Network.MaxConns = 64
	conf.Network.EnableNATService = false
	conf.Network.EnableUPnP = false
	conf.Network.EnableRelay = true
	conf.Network.NetworkName = "aerium"
	conf.Network.DefaultPort = 19933
	conf.Network.DefaultBootstrapAddrStrings = bootstrapAddrs
	conf.GRPC.Enable = true
	conf.GRPC.Listen = "127.0.0.1:50051"
	conf.GRPC.BasicAuth = ""
	conf.HTML.Enable = false
	conf.HTML.Listen = "127.0.0.1:80"
	conf.HTTP.Enable = false
	conf.HTTP.Listen = "127.0.0.1:8080"
	conf.HTTP.BasePath = "/http"
	conf.HTTP.Origins = []string{}
	conf.JSONRPC.Enable = false
	conf.JSONRPC.Listen = "127.0.0.1:8545"
	conf.JSONRPC.Origins = []string{}
	conf.HTML.EnablePprof = false

	return conf
}

func DefaultConfigTestnet() *Config {
	conf := defaultConfig()
	conf.Network.DefaultBootstrapAddrStrings = []string{}
	conf.Network.MaxConns = 64
	conf.Network.EnableNATService = false
	conf.Network.EnableUPnP = false
	conf.Network.EnableRelay = true
	conf.Network.NetworkName = "aerium-testnet"
	conf.Network.DefaultPort = 19944
	conf.GRPC.Enable = true
	conf.GRPC.Listen = "[::]:50052"
	conf.HTML.Enable = false
	conf.HTML.Listen = "[::]:80"
	conf.HTTP.Enable = true
	conf.HTTP.Listen = "[::]:8080"
	conf.HTTP.BasePath = "/http"
	conf.HTTP.Origins = []string{}
	conf.JSONRPC.Enable = true
	conf.JSONRPC.Listen = "[::]:8545"
	conf.JSONRPC.Origins = []string{}
	conf.HTML.EnablePprof = false

	return conf
}

func DefaultConfigLocalnet() *Config {
	conf := defaultConfig()
	conf.Network.EnableRelay = false
	conf.Network.EnableNATService = false
	conf.Network.EnableUPnP = false
	conf.Network.BootstrapAddrStrings = []string{}
	conf.Network.MaxConns = 16
	conf.Network.NetworkName = "aerium-localnet"
	conf.Network.DefaultPort = 0
	conf.Network.ForcePrivateNetwork = true
	conf.Network.EnableMdns = true
	conf.Sync.Moniker = "localnet-1"
	conf.GRPC.Enable = true
	conf.GRPC.EnableWallet = true
	conf.GRPC.Listen = "[::]:50052"
	conf.HTML.Enable = true
	conf.HTML.Listen = "[::]:0"
	conf.HTML.EnablePprof = true
	conf.HTTP.Enable = true
	conf.HTTP.Listen = "[::]:8080"
	conf.HTTP.Origins = []string{"*"}
	conf.JSONRPC.Enable = true
	conf.JSONRPC.Listen = "[::]:8545"
	conf.JSONRPC.Origins = []string{"*"}
	conf.ZeroMq.ZmqPubBlockInfo = "tcp://127.0.0.1:28332"
	conf.ZeroMq.ZmqPubTxInfo = "tcp://127.0.0.1:28333"
	conf.ZeroMq.ZmqPubRawBlock = "tcp://127.0.0.1:28334"
	conf.ZeroMq.ZmqPubRawTx = "tcp://127.0.0.1:28335"
	conf.ZeroMq.ZmqPubHWM = 1000

	return conf
}

func SaveMainnetConfig(path string) error {
	conf := string(exampleConfigBytes)

	return util.WriteFile(path, []byte(conf))
}

func (conf *Config) Save(path string) error {
	return util.WriteFile(path, conf.toTOML())
}

func (conf *Config) toTOML() []byte {
	buf := new(bytes.Buffer)
	encoder := toml.NewEncoder(buf)
	encoder.SetIndentTables(true)
	err := encoder.Encode(conf)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func LoadFromFile(file string, strict bool, defaultConfig *Config) (*Config, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	conf := defaultConfig
	buf := bytes.NewBuffer(data)
	decoder := toml.NewDecoder(buf)
	if strict {
		decoder.DisallowUnknownFields()
	}
	if err := decoder.Decode(conf); err != nil {
		return nil, err
	}

	return conf, nil
}

// BasicCheck performs basic checks on the configuration.
func (conf *Config) BasicCheck() error {
	if err := conf.Node.BasicCheck(); err != nil {
		return err
	}
	if err := conf.Store.BasicCheck(); err != nil {
		return err
	}
	if err := conf.TxPool.BasicCheck(); err != nil {
		return err
	}
	if err := conf.Consensus.BasicCheck(); err != nil {
		return err
	}
	if err := conf.Network.BasicCheck(); err != nil {
		return err
	}
	if err := conf.Logger.BasicCheck(); err != nil {
		return err
	}
	if err := conf.Sync.BasicCheck(); err != nil {
		return err
	}
	if err := conf.JSONRPC.BasicCheck(); err != nil {
		return err
	}
	if err := conf.HTTP.BasicCheck(); err != nil {
		return err
	}
	if err := conf.GRPC.BasicCheck(); err != nil {
		return err
	}
	if err := conf.ZeroMq.BasicCheck(); err != nil {
		return err
	}

	return conf.HTTP.BasicCheck()
}
