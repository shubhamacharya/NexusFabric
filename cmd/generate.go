package cmd

import (
	"fmt"

	"github.com/shubhamacharya/NexusFabric/internal/controllers"
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
			Use:   "manifest",
			Short: "Generate manifest / compose files for the Hyperledger Fabric network.",
			RunE: func(cmd *cobra.Command, args []string) error {
				if netname == "" {
					fmt.Println("Netname not found.... ")
					return nil
				} else {
					fmt.Println("Netname found.... ")
					controllers.BuildAllNetworkComposeFiles(netname)
					return nil
				}
			},
		},
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
