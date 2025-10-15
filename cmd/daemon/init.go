package main

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/aerium-network/aerium/cmd"
	"github.com/aerium-network/aerium/genesis"
	"github.com/aerium-network/aerium/util"
	"github.com/aerium-network/aerium/wallet"
	"github.com/spf13/cobra"
)

// buildInitCmd builds a sub-command to initialize the Aerium blockchain node.
func buildInitCmd(parentCmd *cobra.Command) {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "initialize the Aerium Blockchain node",
	}
	parentCmd.AddCommand(initCmd)

	workingDirOpt := addWorkingDirOption(initCmd)

	testnetOpt := initCmd.Flags().Bool("testnet", false,
		"initialize working directory for joining the testnet")

	localnetOpt := initCmd.Flags().Bool("localnet", false,
		"initialize working directory for localnet (for development)")

	restoreOpt := initCmd.Flags().String("restore", "",
		"restore the 'default_wallet' using a mnemonic or seed phrase")

	passwordOpt := initCmd.Flags().StringP("password", "p", "",
		"the wallet password")

	entropyOpt := initCmd.Flags().IntP("entropy", "e", 128,
		"entropy bits for seed generation. range: 128 to 256")

	valNumOpt := initCmd.Flags().IntP("val-num", "", 0,
		"number of validators to be created. range: 1 to 32")

	initCmd.Run = func(_ *cobra.Command, _ []string) {
		workingDir, _ := filepath.Abs(*workingDirOpt)
		if !util.IsDirNotExistsOrEmpty(workingDir) {
			cmd.PrintErrorMsgf("The working directory is not empty: %s", workingDir)

			return
		}

		index := 0
		recoveryEventFunc := func(addr string) {
			cmd.PrintInfoMsgf("%d. %s", index+1, addr)
			index++
		}

		var mnemonic string
		if *restoreOpt == "" {
			mnemonic, _ = wallet.GenerateMnemonic(*entropyOpt)

			cmd.PrintLine()
			cmd.PrintInfoMsgf("Your wallet seed is:")
			cmd.PrintInfoMsgBoldf("   " + mnemonic)
			cmd.PrintLine()
			cmd.PrintWarnMsgf("Write down this seed on a piece of paper to recover your validator key in the future.")
			cmd.PrintLine()
			confirmed := cmd.PromptConfirm("Do you want to continue")
			if !confirmed {
				return
			}
			recoveryEventFunc = nil
		} else {
			mnemonic = *restoreOpt
			err := wallet.CheckMnemonic(*restoreOpt)
			cmd.FatalErrorCheck(err)
		}

		var password string
		if *passwordOpt == "" {
			cmd.PrintLine()
			cmd.PrintInfoMsgf("Enter a password for wallet")
			password = cmd.PromptPassword("Password", true)
		} else {
			password = *passwordOpt
		}

		var valNum int
		if *valNumOpt == 0 {
			cmd.PrintLine()
			cmd.PrintInfoMsgBoldf("How many validators do you want to create?")
			cmd.PrintInfoMsgf("Each node can run up to 32 validators, and each validator can hold up to 1000 staked coins.")
			cmd.PrintInfoMsgf("You can define validators based on the amount of coins you want to stake.")
			valNum = cmd.PromptInputWithRange("Number of Validators", 7, 1, 32)
		} else {
			if *valNumOpt < 1 || *valNumOpt > 32 {
				cmd.PrintErrorMsgf("%v is not in valid range of validator number, it should be between 1 and 32", *valNumOpt)

				return
			}
			valNum = *valNumOpt
		}

		// TODO: currently genesis hardened to testnet for mainnet we change to genesis.Mainnet.
		chain := genesis.Testnet
		// The order of checking the network (chain type) matters here.
		if *testnetOpt {
			chain = genesis.Testnet
		}
		if *localnetOpt {
			chain = genesis.Localnet
		}

		ctx, cancel := context.WithCancel(context.Background())
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			<-sigChan
			cancel()
		}()

		if recoveryEventFunc != nil {
			cmd.PrintLine()
			cmd.PrintInfoMsgf("Recovering wallet addresses (Ctrl+C to abort)...")
			cmd.PrintLine()
		}

		validatorAddrs, rewardAddrs, err := cmd.CreateNode(ctx, valNum, chain, workingDir, mnemonic,
			password, recoveryEventFunc)
		cmd.FatalErrorCheck(err)

		cmd.PrintLine()
		cmd.PrintInfoMsgBoldf("Validator addresses:")
		for i, addr := range validatorAddrs {
			cmd.PrintInfoMsgf("%v- %s", i+1, addr)
		}
		cmd.PrintLine()

		cmd.PrintInfoMsgBoldf("Reward address:")
		cmd.PrintInfoMsgf("%s", rewardAddrs)

		cmd.PrintLine()
		cmd.PrintInfoMsgBoldf("Network: %v", chain.String())
		cmd.PrintLine()
		cmd.PrintSuccessMsgf("A aerium node is successfully initialized at %v", workingDir)
		cmd.PrintLine()
		cmd.PrintInfoMsgf("You can start the node by running this command:")
		cmd.PrintInfoMsgf("./aerium-daemon start -w %v", workingDir)
	}
}
