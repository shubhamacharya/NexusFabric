package cmd

import (
	"github.com/spf13/cobra"
)

// Command groups for Docker and Kubernetes
var PlatformDockerCmd = NewPlatformCommand("docker", "Setup Hyperledger Fabric network using Docker platform.")
var Platformk8sCmd = NewPlatformCommand("k8s", "Setup Hyperledger Fabric network using Kubernetes platform.")
var netname string

// Common Command Groups
func NewPlatformCommand(name, shortDesc string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   name,
		Short: shortDesc,
	}
	cmd.AddCommand(
		NewGenerateCommand(),
		NewChannelCommand(),
		NewChaincodeCommand(),
	)

	return cmd
}
