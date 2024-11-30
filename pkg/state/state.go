package state

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/shubhamacharya/NexusFabric/internal/constants"
	models "github.com/shubhamacharya/NexusFabric/internal/models"
)

type NetworkState struct {
	Networks map[string]interface{} `json:"networks"`
}

type OrgInfo struct {
	Present bool
	PeerQty int
}

// CreateState generates the network state from the configuration
func CreateState(networkConfig *models.Config) {
	// Initialize the network object
	var network models.Network
	network.Name = networkConfig.Fabric.Netname

	// Assign port based on the configuration
	startPort := networkConfig.Fabric.StartPort
	if startPort == 0 {
		startPort = constants.DefaultPort
	}
	network.Port = models.NewPort(startPort)

	network.CA = models.BuildNewCA("ca."+network.Name, constants.TLSEnabled, network.Port)
	network.CAOrderer = models.BuildNewCA("ca.orderer."+strings.SplitN(networkConfig.Fabric.Orderers[0], ".", 2)[1], constants.TLSEnabled, network.Port)

	for _, orderer := range networkConfig.Fabric.Orderers {
		tempOrderer := models.BuildNewOrderer(orderer, constants.TLSEnabled, network.Port)
		network.Orderers = append(network.Orderers, tempOrderer)
	}
	// Create a map to track organizations and their peer counts
	orgs := make(map[string]OrgInfo)

	// Populate the orgs map with peer information
	for _, peer := range networkConfig.Fabric.Peers {
		org := strings.SplitN(peer, ".", 2)[1] // Extract organization name
		orgInfo := orgs[org]
		orgInfo.PeerQty++
		orgs[org] = orgInfo
	}

	// Build organization structures and add them to the network
	for org, orgInfo := range orgs {
		if !orgInfo.Present {
			orgInfo.Present = true
			orgs[org] = orgInfo

			tempOrg := models.BuildNewOrganization(org, "ca."+org, constants.TLSEnabled, network.Port, orgInfo.PeerQty)
			network.Organizations = append(network.Organizations, tempOrg)
			network.QtyOrgs += 1
		}

	}

	// Serialize the network state to JSON
	state, err := json.MarshalIndent(network, "", " ")
	if err != nil {
		fmt.Printf("Error marshalling network state: %v\n", err)
		return
	}

	networkPaths := models.NewPath(&network)
	// Construct the state file path
	stateDir := filepath.Join("./networks", networkConfig.Fabric.Netname)

	// Ensure the directory exists
	if err := os.MkdirAll(stateDir, 0755); err != nil {
		fmt.Printf("Error creating directories: %v\n", err)
		return
	}

	// Write the state to the file
	if err := os.WriteFile(networkPaths.NETWORKSTATEPATH, state, 0644); err != nil {
		fmt.Printf("Error writing state to file: %v\n", err)
	}
}

func LoadState(filePath string) (*NetworkState, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var state NetworkState
	if err := json.NewDecoder(file).Decode(&state); err != nil {
		return nil, err
	}
	return &state, nil
}

func SaveState(state *models.Network) error {
	networkPath := models.NewPath(state)
	// Ensure the directory exists
	dir := filepath.Dir(networkPath.NETWORKSTATEPATH)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directories for %s: %w", dir, err)
	}

	// Create or overwrite the file
	file, err := os.OpenFile(networkPath.NETWORKSTATEPATH, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to create or open file %s: %w", networkPath.NETWORKSTATEPATH, err)
	}
	defer file.Close()

	// Encode the state as JSON and write to the file
	if err := json.NewEncoder(file).Encode(state); err != nil {
		return fmt.Errorf("failed to encode state to JSON: %w", err)
	}

	return nil
}
