package cmd

import (
	"github.com/spf13/cobra"
)

var ccName, ccPath, ccLang, ccVersion, ccSequence string

func NewChaincodeCommand() *cobra.Command {
	chaincodeCmd := &cobra.Command{
		Use:   "chaincode",
		Short: "Manage chaincode for the Hyperledger Fabric network.",
	}

	// Add subcommands
	chaincodeCmd.AddCommand(
		&cobra.Command{
			Use:   "package",
			Short: "Package chaincode for the Hyperledger Fabric network.",
		},
		&cobra.Command{
			Use:   "install",
			Short: "Install chaincode on peers of the Hyperledger Fabric network.",
		},
		&cobra.Command{
			Use:   "approve",
			Short: "Approve chaincode on the Hyperledger Fabric network.",
		},
		&cobra.Command{
			Use:   "commit",
			Short: "Commit chaincode on the Hyperledger Fabric network.",
		},
	)

	// Add common flags
	chaincodeCmd.PersistentFlags().StringVar(&ccName, "ccName", "", "Name of chaincode to be installed.")
	chaincodeCmd.PersistentFlags().StringVar(&ccPath, "ccPath", "", "Location of the chaincode folder.")
	chaincodeCmd.PersistentFlags().StringVar(&ccLang, "ccLang", "", "Programming language of the chaincode.")
	chaincodeCmd.PersistentFlags().StringVar(&ccVersion, "ccVersion", "", "Version of the chaincode.")
	chaincodeCmd.PersistentFlags().StringVar(&ccSequence, "ccSequence", "", "Sequence of the chaincode.")

	return chaincodeCmd
}
