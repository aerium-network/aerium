package main

import (
	"github.com/aerium-network/aerium/cmd"
	"github.com/aerium-network/aerium/version"
	"github.com/spf13/cobra"
)

func init() {
	version.NodeAgent.AppType = "daemon"
}

func main() {
	rootCmd := &cobra.Command{
		Use:               "aerium-daemon",
		Short:             "Aerium daemon",
		CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
	}

	// Hide the "help" sub-command
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	buildVersionCmd(rootCmd)
	buildInitCmd(rootCmd)
	buildStartCmd(rootCmd)
	buildPruneCmd(rootCmd)
	buildImportCmd(rootCmd)
	buildServiceCmd(rootCmd)

	err := rootCmd.Execute()
	if err != nil {
		cmd.PrintErrorMsgf("%s", err)
	}
}

func addWorkingDirOption(c *cobra.Command) *string {
	return c.Flags().StringP("working-dir", "w", cmd.AeriumDefaultHomeDir(),
		"the path to the working directory that keeps the wallets and node files")
}
