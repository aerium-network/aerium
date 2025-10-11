package genesis_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/aerium-network/aerium/crypto"
	"github.com/aerium-network/aerium/crypto/hash"
	"github.com/aerium-network/aerium/genesis"
	"github.com/aerium-network/aerium/types/account"
	"github.com/aerium-network/aerium/types/amount"
	"github.com/aerium-network/aerium/types/validator"
	"github.com/aerium-network/aerium/util"
	"github.com/aerium-network/aerium/util/testsuite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshaling(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	acc, addr := ts.GenerateTestAccount(
		testsuite.AccountWithNumber(0),
		testsuite.AccountWithBalance(100000))
	val := ts.GenerateTestValidator(
		testsuite.ValidatorWithNumber(0),
	)
	gen1 := genesis.MakeGenesis(util.RoundNow(10),
		map[crypto.Address]*account.Account{addr: acc},
		[]*validator.Validator{val}, genesis.DefaultGenesisParams())
	gen2 := new(genesis.Genesis)

	assert.Equal(t, 10, gen1.Params().BlockIntervalInSecond)

	bz, err := json.MarshalIndent(gen1, " ", " ")
	require.NoError(t, err)
	err = json.Unmarshal(bz, gen2)
	require.NoError(t, err)
	require.Equal(t, gen1.Hash(), gen2.Hash())

	// Test saving and loading
	f := util.TempFilePath()
	assert.NoError(t, gen1.SaveToFile(f))
	gen3, err := genesis.LoadFromFile(f)
	assert.NoError(t, err)
	require.Equal(t, gen1.Hash(), gen3.Hash())

	_, err = genesis.LoadFromFile(util.TempFilePath())
	assert.Error(t, err, "file not found")
}

func TestCheckGenesisAccountAndValidator(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	accs := map[crypto.Address]*account.Account{}
	vals := []*validator.Validator{}
	for i := int32(0); i < 10; i++ {
		pub, _ := ts.RandBLSKeyPair()
		acc := account.NewAccount(i)
		val := validator.NewValidator(pub, i)

		accs[pub.AccountAddress()] = acc
		vals = append(vals, val)
	}
	gen := genesis.MakeGenesis(time.Now(), accs, vals, genesis.DefaultGenesisParams())

	for addr, acc := range gen.Accounts() {
		assert.Equal(t, accs[addr], acc)
	}

	for i, val := range gen.Validators() {
		assert.Equal(t, vals[i].Hash(), val.Hash())
	}
}

func TestGenesisTestnet(t *testing.T) {
	crypto.AddressHRP = "tae"

	gen := genesis.TestnetGenesis()
	fmt.Println(gen.Hash().String())
	assert.Equal(t, 4, len(gen.Validators()))
	assert.Equal(t, 2, len(gen.Accounts()))

	genTime, _ := time.Parse("2006-01-02", "2025-10-15")
	expected, _ := hash.FromString("6ae7151059c16367b4f26e946678c66800267bfbc677bb00a7c6ea4a6429c22d")
	assert.Equal(t, expected, gen.Hash())
	assert.Equal(t, genTime, gen.GenesisTime())
	assert.Equal(t, uint32(360), gen.Params().BondInterval)
	assert.Equal(t, genesis.Testnet, gen.ChainType())
	assert.Equal(t, amount.Amount(amount.TestnetMaxNanoAUM), gen.TotalSupply())
	assert.True(t, gen.ChainType().IsTestnet())

	crypto.AddressHRP = "ae"
}

func TestValidateMainnetSuccess(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	acc, addr := ts.GenerateTestAccount(
		testsuite.AccountWithNumber(0),
		testsuite.AccountWithBalance(amount.DefaultMaxNanoAUM))
	val := ts.GenerateTestValidator(testsuite.ValidatorWithNumber(0))
	gen := genesis.MakeGenesis(
		util.RoundNow(10),
		map[crypto.Address]*account.Account{addr: acc},
		[]*validator.Validator{val},
		genesis.DefaultGenesisParams())

	err := gen.Validate()
	assert.NoError(t, err)
}

func TestValidateMainnetExceedsMaxSupply(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	acc, addr := ts.GenerateTestAccount(
		testsuite.AccountWithNumber(0),
		testsuite.AccountWithBalance(amount.DefaultMaxNanoAUM+1))
	val := ts.GenerateTestValidator(testsuite.ValidatorWithNumber(0))
	gen := genesis.MakeGenesis(
		util.RoundNow(10),
		map[crypto.Address]*account.Account{addr: acc},
		[]*validator.Validator{val},
		genesis.DefaultGenesisParams())

	err := gen.Validate()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "total supply")
	assert.Contains(t, err.Error(), "exceeds maximum")
}

func TestValidateTestnetSuccess(t *testing.T) {
	crypto.AddressHRP = "tae"
	defer func() { crypto.AddressHRP = "ae" }()

	gen := genesis.TestnetGenesis()

	err := gen.Validate()
	assert.NoError(t, err)
}

func TestValidateLocalnetSuccess(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	acc, addr := ts.GenerateTestAccount(
		testsuite.AccountWithNumber(0),
		testsuite.AccountWithBalance(amount.DefaultMaxNanoAUM))
	val := ts.GenerateTestValidator(testsuite.ValidatorWithNumber(0))
	gen := genesis.MakeGenesis(
		time.Now(),
		map[crypto.Address]*account.Account{addr: acc},
		[]*validator.Validator{val},
		genesis.DefaultGenesisParams())

	err := gen.Validate()
	assert.NoError(t, err)
}

func TestValidateLocalnetExceedsMaxSupply(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	acc, addr := ts.GenerateTestAccount(
		testsuite.AccountWithNumber(0),
		testsuite.AccountWithBalance(amount.DefaultMaxNanoAUM+1000))
	val := ts.GenerateTestValidator(testsuite.ValidatorWithNumber(0))
	gen := genesis.MakeGenesis(
		time.Now(),
		map[crypto.Address]*account.Account{addr: acc},
		[]*validator.Validator{val},
		genesis.DefaultGenesisParams())

	err := gen.Validate()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "total supply")
	assert.Contains(t, err.Error(), "exceeds maximum")
	assert.Contains(t, err.Error(), "Localnet")
}

func TestValidateExactMaxSupplyMainnet(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	acc, addr := ts.GenerateTestAccount(
		testsuite.AccountWithNumber(0),
		testsuite.AccountWithBalance(amount.DefaultMaxNanoAUM))
	val := ts.GenerateTestValidator(testsuite.ValidatorWithNumber(0))
	gen := genesis.MakeGenesis(
		util.RoundNow(10),
		map[crypto.Address]*account.Account{addr: acc},
		[]*validator.Validator{val},
		genesis.DefaultGenesisParams())

	err := gen.Validate()
	assert.NoError(t, err)
}

func TestValidateZeroSupply(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	acc, addr := ts.GenerateTestAccount(
		testsuite.AccountWithNumber(0),
		testsuite.AccountWithBalance(0))
	val := ts.GenerateTestValidator(testsuite.ValidatorWithNumber(0))
	gen := genesis.MakeGenesis(
		util.RoundNow(10),
		map[crypto.Address]*account.Account{addr: acc},
		[]*validator.Validator{val},
		genesis.DefaultGenesisParams())

	err := gen.Validate()
	assert.NoError(t, err)
}

func TestValidateMultipleAccountsBelowLimit(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	accs := make(map[crypto.Address]*account.Account)
	for i := 0; i < 5; i++ {
		acc, addr := ts.GenerateTestAccount(
			testsuite.AccountWithNumber(int32(i)),
			testsuite.AccountWithBalance(amount.DefaultMaxNanoAUM/10))
		accs[addr] = acc
	}
	val := ts.GenerateTestValidator(testsuite.ValidatorWithNumber(0))
	gen := genesis.MakeGenesis(
		util.RoundNow(10),
		accs,
		[]*validator.Validator{val},
		genesis.DefaultGenesisParams())

	err := gen.Validate()
	assert.NoError(t, err)
}

func TestValidateMultipleAccountsExceedLimit(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	accs := make(map[crypto.Address]*account.Account)
	for i := 0; i < 5; i++ {
		acc, addr := ts.GenerateTestAccount(
			testsuite.AccountWithNumber(int32(i)),
			testsuite.AccountWithBalance(amount.DefaultMaxNanoAUM/4))
		accs[addr] = acc
	}
	val := ts.GenerateTestValidator(testsuite.ValidatorWithNumber(0))
	gen := genesis.MakeGenesis(
		util.RoundNow(10),
		accs,
		[]*validator.Validator{val},
		genesis.DefaultGenesisParams())

	err := gen.Validate()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "exceeds maximum")
}
