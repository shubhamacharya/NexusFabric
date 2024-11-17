package cmd

import (
	"github.com/spf13/cobra"
)

func NewGenerateCommand() *cobra.Command {
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate artifacts, crypto material, and profiles files.",
	}

	// Add subcommands
	generateCmd.AddCommand(
		&cobra.Command{
			Use:   "genesis",
			Short: "Generate genesis block file for the Hyperledger Fabric network.",
		},
		&cobra.Command{
			Use:   "crypto",
			Short: "Generate crypto material for the Hyperledger Fabric network.",
		},
		&cobra.Command{
			Use:   "profile",
			Short: "Generate connection profile files for the Hyperledger Fabric network.",
		},
	)

	return generateCmd
}
