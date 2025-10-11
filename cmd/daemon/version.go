package main

import (
	"github.com/aerium-network/aerium/version"
	"github.com/spf13/cobra"
)

// Version prints the version of the Aerium node.
func buildVersionCmd(parentCmd *cobra.Command) {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "prints the Aerium version",
	}
	parentCmd.AddCommand(versionCmd)
	versionCmd.Run = func(c *cobra.Command, _ []string) {
		c.Printf("Aerium version: %s\n", version.NodeVersion().StringWithAlias())
	}
}
