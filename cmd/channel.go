package cmd

import (
	"github.com/spf13/cobra"
)

func NewChannelCommand() *cobra.Command {
	channelCmd := &cobra.Command{
		Use:   "channel",
		Short: "Manage channels for the Hyperledger Fabric network.",
	}

	// Add subcommands
	channelCmd.AddCommand(
		&cobra.Command{
			Use:   "create",
			Short: "Create new channels for the Hyperledger Fabric network.",
		},
		&cobra.Command{
			Use:   "join",
			Short: "Join peers to channels for the Hyperledger Fabric network.",
		},
		&cobra.Command{
			Use:   "anchorUpdate",
			Short: "Update anchor peers for organizations in the Hyperledger Fabric network.",
		},
	)

	return channelCmd
}
