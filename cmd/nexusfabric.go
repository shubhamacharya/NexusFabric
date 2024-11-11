package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nexusfabric",
	Short: "NexusFabric CLI to manage Hyperledger Fabric networks on Docker and Kubernetes",
}

rootCmd.AddCommand()

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
