package cmd

import (
	"github.com/spf13/cobra"
)

var Platformk8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Setup hyperledger fabric network using k8s platform.",
}

var Generatek8sCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate artifacts, cryto material and profiles files.",
}
var GenerateGenesisk8sCmd = &cobra.Command{
	Use:   "genesis",
	Short: "Generate genesis block file for Hyperledger fabric network.",
}
var GenerateCryptok8sCmd = &cobra.Command{
	Use:   "crypto",
	Short: "Generate crytpo material for Hyperledger fabric network.",
}
var GenerateProfilesk8sCmd = &cobra.Command{
	Use:   "profile",
	Short: "Generate Connection Profile files for Hyperledger fabric network.",
}

var Channelk8sCmd = &cobra.Command{
	Use:   "channel",
	Short: "Create and Join the channels with Anchor Update.",
}
var ChannelCreatek8sCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates new channel / channels for Hyperledger fabric network.",
}
var ChannelJoink8sCmd = &cobra.Command{
	Use:   "join",
	Short: "Joins peer / peers to new channel / channels.",
}
var ChannelAchorUpdatek8sCmd = &cobra.Command{
	Use:   "anchorUpdate",
	Short: "Updated the anchor peers for Organizations of Hyperledger fabric network.",
}

var Chaincodek8sCmd = &cobra.Command{
	Use:   "chaincode",
	Short: "Package, Install, Approve and Commit the chaincode on Hyperledger Fabric Network.",
}
var PackageChaincodek8sCmd = &cobra.Command{
	Use:   "package",
	Short: "Packages the chaincode for Hyperledger fabric network.",
}
var InstallChaincodek8sCmd = &cobra.Command{
	Use:   "install",
	Short: "Install chaincode on peers of Hyperledger fabric network.",
}
var ApproveChaincodek8sCmd = &cobra.Command{
	Use:   "approve",
	Short: "Approve chaincode on Hyperledger fabric network.",
}
var CommitChaincodek8sCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit chaincode on Hyperledger fabric network.",
}

func init() {
	Generatek8sCmd.AddCommand(GenerateGenesisk8sCmd, GenerateCryptok8sCmd, GenerateProfilesk8sCmd)
	Channelk8sCmd.AddCommand(ChannelCreatek8sCmd, ChannelJoink8sCmd, ChannelAchorUpdatek8sCmd)
	Chaincodek8sCmd.AddCommand(PackageChaincodek8sCmd, InstallChaincodek8sCmd, ApproveChaincodek8sCmd, CommitChaincodek8sCmd)

	Platformk8sCmd.AddCommand(Generatek8sCmd, Channelk8sCmd, Chaincodek8sCmd)
}
