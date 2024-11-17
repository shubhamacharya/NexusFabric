package controllers

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	models "github.com/shubhamacharya/NexusFabric/internal/models"
	state "github.com/shubhamacharya/NexusFabric/pkg/state"
	"gopkg.in/yaml.v3"
)

func ReadSpecsFile(networkConfigFilePath string) {
	if networkConfigFilePath == "" {
		networkConfigFilePath, _ = filepath.Abs(".")
		networkConfigFilePath += "/network_config.yaml"
	}

	networkConfigFile, err := os.ReadFile(networkConfigFilePath)
	if err != nil {
		fmt.Printf("Error reading Network Config File: %v\n", err)
		return
	}

	var config models.Config
	err = yaml.Unmarshal(networkConfigFile, &config)
	if err != nil {
		fmt.Printf("Error unmarshaling YAML: %v\n", err)
		return
	}
	state.CreateState(&config)

	// fmt.Printf("Fabric Config:\n%+v\n", config.Fabric)
	// fmt.Printf("CAs: %v\n", config.Fabric.CAs)
	// fmt.Printf("Peers: %v\n", config.Fabric.Peers)
	// fmt.Printf("Orderers: %v\n", config.Fabric.Orderers)
	// fmt.Printf("Netname: %s\n", config.Fabric.Netname)
	// fmt.Printf("CA Logging Spec: %s\n", config.Fabric.Settings.CA.FabricLoggingSpec)
}

func CheckHLFBinaries(network *models.Network, paths *models.Paths) error {
	if _, err := os.Stat(paths.NETWORKPATH); err == nil {
		fmt.Println("File Exists")
	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist

	} else {
		// Schrodinger: file may or may not exist. See err for details.

		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence

	}

	return nil
}
