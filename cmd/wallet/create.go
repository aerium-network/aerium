package main

import (
	"github.com/aerium-network/aerium/cmd"
	"github.com/aerium-network/aerium/genesis"
	"github.com/aerium-network/aerium/wallet"
	"github.com/spf13/cobra"
)

// buildCreateCmd builds a command to create a new wallet.
func buildCreateCmd(parentCmd *cobra.Command) {
	generateCmd := &cobra.Command{
		Use:   "create",
		Short: "creating a new wallet",
	}
	parentCmd.AddCommand(generateCmd)

	testnetOpt := generateCmd.Flags().Bool("testnet", false,
		"create a wallet for the testnet environment")
	entropyOpt := generateCmd.Flags().Int("entropy", 256,
		"specify the entropy bit length")

	generateCmd.Run = func(_ *cobra.Command, _ []string) {
		password := cmd.PromptPassword("Password", true)
		mnemonic, err := wallet.GenerateMnemonic(*entropyOpt)
		cmd.FatalErrorCheck(err)

		// TODO: currently genesis hardened to testnet for mainnet we change to genesis.Mainnet.
		network := genesis.Testnet
		if *testnetOpt {
			network = genesis.Testnet
		}
		wlt, err := wallet.Create(*pathOpt, mnemonic, password, network)
		cmd.FatalErrorCheck(err)

		err = wlt.Save()
		cmd.FatalErrorCheck(err)

		cmd.PrintLine()
		cmd.PrintSuccessMsgf("Your wallet was successfully created at: %s", wlt.Path())
		cmd.PrintInfoMsgf("Seed phrase: \"%v\"", mnemonic)
		cmd.PrintWarnMsgf("Please keep your seed in a safe place; " +
			"if you lose it, you will not be able to restore your wallet.")
	}
}

// buildChangePasswordCmd builds a command to update the wallet's password.
func buildChangePasswordCmd(parentCmd *cobra.Command) {
	changePasswordCmd := &cobra.Command{
		Use:   "password",
		Short: "changes the wallet's password",
	}
	parentCmd.AddCommand(changePasswordCmd)
	passOpt := addPasswordOption(changePasswordCmd)

	changePasswordCmd.Run = func(_ *cobra.Command, _ []string) {
		wlt, err := openWallet()
		cmd.FatalErrorCheck(err)

		oldPassword := getPassword(wlt, *passOpt)
		newPassword := cmd.PromptPassword("New Password", true)

		err = wlt.UpdatePassword(oldPassword, newPassword)
		cmd.FatalErrorCheck(err)

		err = wlt.Save()
		cmd.FatalErrorCheck(err)

		cmd.PrintLine()
		cmd.PrintWarnMsgf("Your wallet password successfully updated.")
	}
}
