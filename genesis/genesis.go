package genesis

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aerium-network/aerium/crypto"
	"github.com/aerium-network/aerium/crypto/bls"
	"github.com/aerium-network/aerium/crypto/hash"
	"github.com/aerium-network/aerium/types/account"
	"github.com/aerium-network/aerium/types/amount"
	"github.com/aerium-network/aerium/types/validator"
	"github.com/aerium-network/aerium/util"
	"github.com/fxamacker/cbor/v2"
)

type ChainType uint8

const (
	Mainnet  ChainType = 0
	Testnet  ChainType = 1
	Localnet ChainType = 2
)

func (n ChainType) IsMainnet() bool {
	return n == Mainnet
}

func (n ChainType) IsTestnet() bool {
	return n == Testnet
}

func (n ChainType) String() string {
	switch n {
	case Mainnet:
		return "Mainnet"
	case Testnet:
		return "Testnet"
	case Localnet:
		return "Localnet"
	default:
		return "Unknown"
	}
}

type genAccount struct {
	Address string        `cbor:"1,keyasint" json:"address"`
	Balance amount.Amount `cbor:"2,keyasint" json:"balance"`
}

type genValidator struct {
	PublicKey string `cbor:"1,keyasint" json:"public_key"`
}

// Genesis is stored in the state database.
type Genesis struct {
	data genesisData
}

type genesisData struct {
	GenesisTime time.Time      `cbor:"1,keyasint" json:"genesis_time"`
	Params      *GenesisParams `cbor:"2,keyasint" json:"params"`
	Accounts    []genAccount   `cbor:"3,keyasint" json:"accounts"`
	Validators  []genValidator `cbor:"4,keyasint" json:"validators"`
}

func (gen *Genesis) Hash() hash.Hash {
	bs, _ := cbor.Marshal(gen.data)

	return hash.CalcHash(bs)
}

func (gen *Genesis) GenesisTime() time.Time {
	return gen.data.GenesisTime
}

func (gen *Genesis) Params() *GenesisParams {
	return gen.data.Params
}

func (gen *Genesis) Accounts() map[crypto.Address]*account.Account {
	accs := make(map[crypto.Address]*account.Account)
	for i, genAcc := range gen.data.Accounts {
		addr, err := crypto.AddressFromString(genAcc.Address)
		if err != nil {
			panic(err)
		}
		acc := account.NewAccount(int32(i))
		acc.AddToBalance(genAcc.Balance)
		accs[addr] = acc
	}

	return accs
}

func (gen *Genesis) Validators() []*validator.Validator {
	vals := make([]*validator.Validator, 0, len(gen.data.Validators))
	for i, genVal := range gen.data.Validators {
		pub, _ := bls.PublicKeyFromString(genVal.PublicKey)
		val := validator.NewValidator(pub, int32(i))
		vals = append(vals, val)
	}

	return vals
}

func (gen *Genesis) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(&gen.data, "  ", "  ")
}

func (gen *Genesis) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &gen.data)
}

func (gen *Genesis) Validate() error {
	chainType := gen.ChainType()
	var maxSupply amount.Amount

	if chainType == Testnet {
		maxSupply = amount.TestnetMaxNanoAUM
	} else {
		maxSupply = amount.DefaultMaxNanoAUM
	}

	totalSupply := gen.TotalSupply()
	if totalSupply > maxSupply {
		return fmt.Errorf("total supply %v exceeds maximum %v for %s",
			totalSupply, maxSupply, chainType.String())
	}

	return nil
}

func makeGenesisAccount(addr crypto.Address, acc *account.Account) genAccount {
	return genAccount{
		Address: addr.String(),
		Balance: acc.Balance(),
	}
}

func makeGenesisValidator(val *validator.Validator) genValidator {
	return genValidator{
		PublicKey: val.PublicKey().String(),
	}
}

func MakeGenesis(genesisTime time.Time, accounts map[crypto.Address]*account.Account,
	validators []*validator.Validator, params *GenesisParams,
) *Genesis {
	genAccs := make([]genAccount, len(accounts))
	for addr, acc := range accounts {
		genAcc := makeGenesisAccount(addr, acc)
		genAccs[acc.Number()] = genAcc
	}

	genVals := make([]genValidator, len(validators))
	for _, val := range validators {
		genVal := makeGenesisValidator(val)
		genVals[val.Number()] = genVal
	}

	return &Genesis{
		data: genesisData{
			GenesisTime: genesisTime,
			Accounts:    genAccs,
			Validators:  genVals,
			Params:      params,
		},
	}
}

// LoadFromFile loads genesis object from a JSON file.
func LoadFromFile(file string) (*Genesis, error) {
	dat, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var gen Genesis
	if err := json.Unmarshal(dat, &gen); err != nil {
		return nil, err
	}

	return &gen, nil
}

// SaveToFile saves the genesis into a JSON file.
func (gen *Genesis) SaveToFile(file string) error {
	data, err := gen.MarshalJSON()
	if err != nil {
		return err
	}

	// write  dataContent to file
	return util.WriteFile(file, data)
}

func (gen *Genesis) TotalSupply() amount.Amount {
	totalSuppyly := amount.Amount(0)
	for _, acc := range gen.data.Accounts {
		totalSuppyly += acc.Balance
	}

	return totalSuppyly
}

func (gen *Genesis) ChainType() ChainType {
	switch gen.Hash() {
	case MainnetGenesis().Hash():
		return Mainnet
	case TestnetGenesis().Hash():
		return Testnet
	default:
		return Localnet
	}
}
