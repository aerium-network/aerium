package cmd

import (
	"bytes"
	"context"
	"io"
	"os"
	"runtime"
	"testing"

	"github.com/aerium-network/aerium/config"
	"github.com/aerium-network/aerium/genesis"
	"github.com/aerium-network/aerium/util"
	"github.com/aerium-network/aerium/util/testsuite"
	"github.com/aerium-network/aerium/wallet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMakeConfig(t *testing.T) {
	t.Run("No genesis file, Should return error", func(t *testing.T) {
		workingDir := util.TempDirPath()

		_, _, err := MakeConfig(workingDir)
		assert.Error(t, err)
	})

	t.Run("No Config file, Should recover it", func(t *testing.T) {
		workingDir := util.TempDirPath()
		genPath := AeriumGenesisPath(workingDir)
		gen := genesis.MainnetGenesis()
		err := gen.SaveToFile(genPath)
		require.NoError(t, err)

		_, _, err = MakeConfig(workingDir)
		assert.NoError(t, err)
	})

	t.Run("Invalid Config file, Should recover it", func(t *testing.T) {
		workingDir := util.TempDirPath()
		genPath := AeriumGenesisPath(workingDir)
		confPath := AeriumConfigPath(workingDir)

		gen := genesis.MainnetGenesis()
		err := gen.SaveToFile(genPath)
		require.NoError(t, err)

		err = util.WriteFile(confPath, []byte("invalid-config"))
		require.NoError(t, err)

		_, _, err = MakeConfig(workingDir)
		assert.NoError(t, err)
	})

	t.Run("Everything is good", func(t *testing.T) {
		workingDir := util.TempDirPath()
		genPath := AeriumGenesisPath(workingDir)
		confPath := AeriumConfigPath(workingDir)

		gen := genesis.MainnetGenesis()
		err := gen.SaveToFile(genPath)
		require.NoError(t, err)

		err = config.SaveMainnetConfig(confPath)
		require.NoError(t, err)

		_, _, err = MakeConfig(workingDir)
		assert.NoError(t, err)
	})
}

// captureOutput is a helper function to capture the printed output of a function.
func captureOutput(fun func()) string {
	// Redirect stdout to a buffer
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	// Capture the printed output
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, reader)
		outC <- buf.String()
	}()

	// Execute the function
	fun()

	// Reset stdout
	_ = writer.Close()
	os.Stdout = oldStdout
	out := <-outC

	return out
}

func TestPrintNotSupported(t *testing.T) {
	terminalSupported = false
	output := captureOutput(func() {
		PrintJSONObject([]int{1, 2, 3})
		PrintLine()
		PrintInfoMsgBoldf("This is PrintInfoMsgBoldf: %s", "msg")
		PrintInfoMsgf("This is PrintInfoMsgf: %s", "msg")
		PrintSuccessMsgf("This is PrintSuccessMsgf: %s", "msg")
		PrintWarnMsgf("This is PrintWarnMsgf: %s", "msg")
		PrintErrorMsgf("This is PrintErrorMsgf: %s", "msg")
	})

	expected := "[\n   1,\n   2,\n   3\n]\n" +
		"\n" +
		"This is PrintInfoMsgBoldf: msg\n" +
		"This is PrintInfoMsgf: msg\n" +
		"This is PrintSuccessMsgf: msg\n" +
		"This is PrintWarnMsgf: msg\n" +
		"[ERROR] This is PrintErrorMsgf: msg\n"

	assert.Equal(t, expected, output)
}

func TestPrintSupported(t *testing.T) {
	terminalSupported = true
	output := captureOutput(func() {
		PrintJSONObject([]int{1, 2, 3})
		PrintLine()
		PrintInfoMsgBoldf("This is PrintInfoMsgBoldf: %s", "msg")
		PrintInfoMsgf("This is PrintInfoMsgf: %s", "msg")
		PrintSuccessMsgf("This is PrintSuccessMsgf: %s", "msg")
		PrintWarnMsgf("This is PrintWarnMsgf: %s", "msg")
		PrintErrorMsgf("This is PrintErrorMsgf: %s", "msg")
	})

	expected := "[\n   1,\n   2,\n   3\n]\n" +
		"\n" +
		"\x1b[1mThis is PrintInfoMsgBoldf: msg\x1b[0m\n" +
		"This is PrintInfoMsgf: msg\n" +
		"\x1b[32mThis is PrintSuccessMsgf: msg\x1b[0m\n" +
		"\x1b[33mThis is PrintWarnMsgf: msg\x1b[0m\n" +
		"\x1b[31m[ERROR] This is PrintErrorMsgf: msg\x1b[0m\n"

	assert.Equal(t, expected, output)
}

func TestPathsUnix(t *testing.T) {
	if runtime.GOOS == "windows" {
		return
	}
	tests := []struct {
		home                      string
		expectedWalletDir         string
		expectedDefaultWalletPath string
		expectedGenesisPath       string
		expectedConfigPath        string
	}{
		{
			"/home/aerium",
			"/home/aerium/wallets",
			"/home/aerium/wallets/default_wallet",
			"/home/aerium/genesis.json",
			"/home/aerium/config.toml",
		},
		{
			"/home/aerium/",
			"/home/aerium/wallets",
			"/home/aerium/wallets/default_wallet",
			"/home/aerium/genesis.json",
			"/home/aerium/config.toml",
		},
	}

	for _, tt := range tests {
		walletDir := AeriumWalletDir(tt.home)
		defaultWalletPath := AeriumDefaultWalletPath(tt.home)
		genesisPath := AeriumGenesisPath(tt.home)
		configPath := AeriumConfigPath(tt.home)

		assert.Equal(t, tt.expectedWalletDir, walletDir)
		assert.Equal(t, tt.expectedDefaultWalletPath, defaultWalletPath)
		assert.Equal(t, tt.expectedGenesisPath, genesisPath)
		assert.Equal(t, tt.expectedConfigPath, configPath)
	}
}

func TestPathsWindows(t *testing.T) {
	if runtime.GOOS != "windows" {
		return
	}
	tests := []struct {
		home                      string
		expectedWalletDir         string
		expectedDefaultWalletPath string
		expectedGenesisPath       string
		expectedConfigPath        string
	}{
		{
			"c:\\aerium",
			"c:\\aerium\\wallets",
			"c:\\aerium\\wallets\\default_wallet",
			"c:\\aerium\\genesis.json",
			"c:\\aerium\\config.toml",
		},
		{
			"c:\\home\\",
			"c:\\home\\wallets",
			"c:\\home\\wallets\\default_wallet",
			"c:\\home\\genesis.json",
			"c:\\home\\config.toml",
		},
	}

	for _, tt := range tests {
		walletDir := AeriumWalletDir(tt.home)
		defaultWalletPath := AeriumDefaultWalletPath(tt.home)
		genesisPath := AeriumGenesisPath(tt.home)
		configPath := AeriumConfigPath(tt.home)

		assert.Equal(t, tt.expectedWalletDir, walletDir)
		assert.Equal(t, tt.expectedDefaultWalletPath, defaultWalletPath)
		assert.Equal(t, tt.expectedGenesisPath, genesisPath)
		assert.Equal(t, tt.expectedConfigPath, configPath)
	}
}

func TestMakeRewardAddresses(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	setupWallet := func() *wallet.Wallet {
		walletPath := util.TempFilePath()
		mnemonic, _ := wallet.GenerateMnemonic(128)
		wlt, err := wallet.Create(walletPath, mnemonic, "", genesis.Mainnet)
		assert.NoError(t, err)

		_, _ = wlt.NewValidatorAddress("Validator 1")
		_, _ = wlt.NewValidatorAddress("Validator 2")
		_, _ = wlt.NewValidatorAddress("Validator 3")

		return wlt
	}

	t.Run("No reward addresses in wallet", func(t *testing.T) {
		wlt := setupWallet()

		valAddrsInfo := wlt.AllValidatorAddresses()
		confRewardAddresses := []string{}
		_, err := MakeRewardAddresses(wlt, valAddrsInfo, confRewardAddresses)
		assert.ErrorContains(t, err, "unable to find a reward address in the wallet")
	})

	t.Run("Wallet with one Ed25519 address", func(t *testing.T) {
		wlt := setupWallet()

		addr1Info, _ := wlt.NewEd25519AccountAddress("", "")
		_, _ = wlt.NewEd25519AccountAddress("", "")
		_, _ = wlt.NewBLSAccountAddress("")

		valAddrsInfo := wlt.AllValidatorAddresses()
		confRewardAddresses := []string{}
		rewardAddrs, err := MakeRewardAddresses(wlt, valAddrsInfo, confRewardAddresses)
		assert.NoError(t, err)

		assert.Equal(t, rewardAddrs[0].String(), addr1Info.Address)
		assert.Equal(t, rewardAddrs[1].String(), addr1Info.Address)
		assert.Equal(t, rewardAddrs[2].String(), addr1Info.Address)
	})

	t.Run("Wallet with one BLS address", func(t *testing.T) {
		wlt := setupWallet()

		addr1Info, _ := wlt.NewBLSAccountAddress("")
		_, _ = wlt.NewBLSAccountAddress("")

		valAddrsInfo := wlt.AllValidatorAddresses()
		confRewardAddresses := []string{}
		rewardAddrs, err := MakeRewardAddresses(wlt, valAddrsInfo, confRewardAddresses)
		assert.NoError(t, err)

		assert.Equal(t, rewardAddrs[0].String(), addr1Info.Address)
		assert.Equal(t, rewardAddrs[1].String(), addr1Info.Address)
		assert.Equal(t, rewardAddrs[2].String(), addr1Info.Address)
	})

	t.Run("One reward address in config", func(t *testing.T) {
		wlt := setupWallet()

		valAddrsInfo := wlt.AllValidatorAddresses()
		confRewardAddresses := []string{
			ts.RandAccAddress().String(),
		}
		rewardAddrs, err := MakeRewardAddresses(wlt, valAddrsInfo, confRewardAddresses)
		assert.NoError(t, err)

		assert.Equal(t, rewardAddrs[0].String(), confRewardAddresses[0])
		assert.Equal(t, rewardAddrs[1].String(), confRewardAddresses[0])
		assert.Equal(t, rewardAddrs[2].String(), confRewardAddresses[0])
	})

	t.Run("Three reward addresses in config", func(t *testing.T) {
		wlt := setupWallet()

		valAddrsInfo := wlt.AllValidatorAddresses()
		confRewardAddresses := []string{
			ts.RandAccAddress().String(),
			ts.RandAccAddress().String(),
			ts.RandAccAddress().String(),
		}
		rewardAddrs, err := MakeRewardAddresses(wlt, valAddrsInfo, confRewardAddresses)
		assert.NoError(t, err)

		assert.Equal(t, rewardAddrs[0].String(), confRewardAddresses[0])
		assert.Equal(t, rewardAddrs[1].String(), confRewardAddresses[1])
		assert.Equal(t, rewardAddrs[2].String(), confRewardAddresses[2])
	})

	t.Run("Insufficient reward addresses in config", func(t *testing.T) {
		wlt := setupWallet()

		valAddrsInfo := wlt.AllValidatorAddresses()
		confRewardAddresses := []string{
			ts.RandAccAddress().String(),
			ts.RandAccAddress().String(),
		}
		_, err := MakeRewardAddresses(wlt, valAddrsInfo, confRewardAddresses)
		assert.ErrorContains(t, err, "expected 3 reward addresses, but got 2")
	})
}

func TestCreateNode(t *testing.T) {
	tests := []struct {
		name           string
		numValidators  int
		chain          genesis.ChainType
		workingDir     string
		mnemonic       string
		withErr        bool
		validatorAddrs []string
		rewardAddrs    string
	}{
		{
			name:           "Create node for Mainnet",
			numValidators:  1,
			chain:          genesis.Mainnet,
			workingDir:     util.TempDirPath(),
			mnemonic:       "legal winner thank year wave sausage worth useful legal winner thank yellow",
			validatorAddrs: []string{"ae1p0lnw2tj2xwasmdltha6wghzk09m2kqa25hl7kn"},
			rewardAddrs:    "ae1r3ra20jvrc9498qr95525awlkur4v42q0wda78x",
			withErr:        false,
		},
		{
			name:           "Create node for Testnet",
			numValidators:  1,
			chain:          genesis.Testnet,
			workingDir:     util.TempDirPath(),
			mnemonic:       "legal winner thank year wave sausage worth useful legal winner thank yellow",
			validatorAddrs: []string{"tae1pqtkxy9rn8kvdk25dljl9mpl56c6puhpg0qc635"},
			rewardAddrs:    "tae1r5qxqgrn8qgrm76lv9yrxass4czvcu5djevr5dq",
			withErr:        false,
		},

		{
			name:          "Create node for Localnet",
			numValidators: 4,
			chain:         genesis.Localnet,
			workingDir:    util.TempDirPath(),
			mnemonic:      "legal winner thank year wave sausage worth useful legal winner thank yellow",
			validatorAddrs: []string{
				"tae1pqtkxy9rn8kvdk25dljl9mpl56c6puhpg0qc635",
				"tae1p2wqn6m5whguasevh4narn6qsy85egdk9z7ugem",
				"tae1pjy9v7fkam3jxsf036atza75z5r5ynmd276nldk",
				"tae1pkdtn3257xg0f6lv0qncak3hys5kw4mygsq04ex",
			},
			rewardAddrs: "tae1r5qxqgrn8qgrm76lv9yrxass4czvcu5djevr5dq",
			withErr:     false,
		},
		{
			name:           "Localnet with one validator",
			numValidators:  1,
			chain:          genesis.Localnet,
			workingDir:     util.TempDirPath(),
			mnemonic:       "legal winner thank year wave sausage worth useful legal winner thank yellow",
			validatorAddrs: nil,
			rewardAddrs:    "",
			withErr:        true,
		},
		{
			name:           "Invalid mnemonic",
			numValidators:  4,
			chain:          genesis.Mainnet,
			workingDir:     util.TempDirPath(),
			mnemonic:       "",
			validatorAddrs: nil,
			rewardAddrs:    "",
			withErr:        true,
		},
	}

	for _, tt := range tests {
		validatorAddrs, rewardAddrs, err := CreateNode(context.Background(),
			tt.numValidators, tt.chain, tt.workingDir, tt.mnemonic, "", nil)

		if tt.withErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tt.validatorAddrs, validatorAddrs)
			assert.Equal(t, tt.rewardAddrs, rewardAddrs)
		}
	}
}
