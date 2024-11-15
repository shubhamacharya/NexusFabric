package cmd

import (
	"github.com/spf13/cobra"
)

var PlatformDockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Setup hyperledger fabric network using Docker platform.",
}

var GenerateDockerCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate artifacts, cryto material and profiles files.",
}
var GenerateGenesisDockerCmd = &cobra.Command{
	Use:   "genesis",
	Short: "Generate genesis block file for Hyperledger fabric network.",
}
var GenerateCryptoDockerCmd = &cobra.Command{
	Use:   "crypto",
	Short: "Generate crytpo material for Hyperledger fabric network.",
}
var GenerateProfilesDockerCmd = &cobra.Command{
	Use:   "profile",
	Short: "Generate Connection Profile files for Hyperledger fabric network.",
}

var ChannelDockerCmd = &cobra.Command{
	Use:   "channel",
	Short: "Create and Join the channels with Anchor Update.",
}
var ChannelCreateDockerCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates new channel / channels for Hyperledger fabric network.",
}
var ChannelJoinDockerCmd = &cobra.Command{
	Use:   "join",
	Short: "Joins peer / peers to new channel / channels.",
}
var ChannelAchorUpdateDockerCmd = &cobra.Command{
	Use:   "anchorUpdate",
	Short: "Updated the anchor peers for Organizations of Hyperledger fabric network.",
}

var ChaincodeDockerCmd = &cobra.Command{
	Use:   "chaincode",
	Short: "Package, Install, Approve and Commit the chaincode on Hyperledger Fabric Network.",
}
var PackageChaincodeDockerCmd = &cobra.Command{
	Use:   "package",
	Short: "Packages the chaincode for Hyperledger fabric network.",
}
var InstallChaincodeDockerCmd = &cobra.Command{
	Use:   "install",
	Short: "Install chaincode on peers of Hyperledger fabric network.",
}
var ApproveChaincodeDockerCmd = &cobra.Command{
	Use:   "approve",
	Short: "Approve chaincode on Hyperledger fabric network.",
}
var CommitChaincodeDockerCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit chaincode on Hyperledger fabric network.",
}

func init() {
	GenerateDockerCmd.AddCommand(GenerateGenesisDockerCmd, GenerateCryptoDockerCmd, GenerateProfilesDockerCmd)
	ChannelDockerCmd.AddCommand(ChannelCreateDockerCmd, ChannelJoinDockerCmd, ChannelAchorUpdateDockerCmd)
	ChaincodeDockerCmd.AddCommand(PackageChaincodeDockerCmd, InstallChaincodeDockerCmd, ApproveChaincodeDockerCmd, CommitChaincodeDockerCmd)

	PlatformDockerCmd.AddCommand(GenerateDockerCmd, ChannelDockerCmd, ChaincodeDockerCmd)
}
