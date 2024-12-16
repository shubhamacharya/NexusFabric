package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "nexusfabric",
	Short: "NexusFabric CLI to manage Hyperledger Fabric networks on Docker and Kubernetes.",
}

func init() {
	RootCmd.AddCommand(PlatformDockerCmd, Platformk8sCmd)
	RootCmd.PersistentFlags().StringVar(&netname, "netName", "", "Name of the network for / on which operation to be performed.")

}
