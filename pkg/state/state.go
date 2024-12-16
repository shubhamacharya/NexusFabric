package state

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/shubhamacharya/NexusFabric/internal/constants"
	models "github.com/shubhamacharya/NexusFabric/internal/models"
)

type NetworkState struct {
	Networks map[string]string `json:"networks"`
}

type OrgInfo struct {
	Present bool
	PeerQty int
}

// CreateState generates the network state from the configuration
func CreateState(networkConfig *models.Config) error {
	// Initialize the network object
	var network models.Network
	network.Name = networkConfig.Fabric.Netname

	// Assign port based on the configuration
	startPort := networkConfig.Fabric.StartPort
	if startPort == 0 {
		startPort = constants.DefaultPort
	}
	network.Port = models.NewPort(startPort)

	// Optionally set Registry URL
	if networkConfig.Fabric.RegistryURL != "" {
		network.RegistryURL = networkConfig.Fabric.RegistryURL
	}

	// Set default images if not provided
	if networkConfig.Fabric.Images == nil {
		network.Images = make(map[string]string)
		network.Images["ca"] = "hyperledger/fabric-ca:" + networkConfig.Fabric.FabricCAVersion
	}

	// Initialize the Certificate Authorities (CA)
	network.CA = models.BuildNewCA("ca."+network.Name, constants.TLSEnabled, network.Port)
	network.CAOrderer = models.BuildNewCA("ca.orderer."+strings.SplitN(networkConfig.Fabric.Orderers[0], ".", 2)[1], constants.TLSEnabled, network.Port)

	// Initialize orderers
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

	// Initialize paths and create directories
	networkPaths := models.NewPath(&network)
	models.CreateNewPaths(&network)

	// Serialize the network state to JSON with indentation
	state, err := json.MarshalIndent(network, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling network state: %v\n", err)
		return err
	}

	// Open the state file for writing the JSON state
	stateFile, err := os.OpenFile(networkPaths.NETWORKSTATEPATH, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to create or open file %s: %w", networkPaths.NETWORKSTATEPATH, err)
	}
	defer stateFile.Close()

	// Write the JSON state to the file
	if _, err := stateFile.Write(state); err != nil {
		return fmt.Errorf("failed to write state to file %s: %w", networkPaths.NETWORKSTATEPATH, err)
	}

	// Construct the state file path
	stateDir, _ := filepath.Abs(filepath.Join("./networks", networkConfig.Fabric.Netname))
	// Create a state list for the network
	stateList := make(map[string]string)
	stateList[network.Name] = networkPaths.NETWORKSTATEPATH

	// Open the state list file to record the network paths
	stateListFile, err := os.OpenFile(networkPaths.NETWORKSTATELISTPATH, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to create or open file %s: %w", networkPaths.NETWORKSTATELISTPATH, err)
	}
	defer stateListFile.Close()

	// Ensure the directory exists
	if err := os.MkdirAll(stateDir, 0755); err != nil {
		fmt.Printf("Error creating directories: %v\n", err)
		return err
	}

	// Write the state list to the file
	stateListJSON, err := json.MarshalIndent(stateList, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal state list to JSON: %w", err)
	}

	if _, err := stateListFile.Write(stateListJSON); err != nil {
		return fmt.Errorf("failed to write state list to file %s: %w", networkPaths.NETWORKSTATELISTPATH, err)
	}

	return nil
}

func LoadStateList() (*NetworkState, error) {
	// Get the path to the network state list (assuming NewPath and NETWORKSTATELISTPATH are correct)
	networkPaths := models.NewPath(new(models.Network))
	file, err := os.Open(networkPaths.NETWORKSTATELISTPATH)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err) // Wrap the error for better context
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	if len(fileContents) == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	// Create an empty NetworkState to load the data into
	var stateList NetworkState

	// Use json.NewDecoder to decode the file directly into stateList
	if err := json.Unmarshal(fileContents, &stateList.Networks); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	// Return the pointer to the decoded NetworkState object
	return &stateList, nil
}

// GetNetworkStateByName retrieves the network state by its name
func GetNetworkStateByName(networkName string) (*models.Network, error) {
	// Load the state list (this assumes LoadStateList returns a pointer)
	stateListPtr, err := LoadStateList()
	if err != nil {
		// If there's an error loading the state list, return nil and the error
		return nil, fmt.Errorf("failed to load state list: %w", err)
	}

	// Dereference the pointer to access the value
	stateList := *stateListPtr
	// Check if the networkName exists in the stateList
	stateFilePath, exists := stateList.Networks[networkName]
	if !exists {
		return nil, fmt.Errorf("network %s not found in state list", networkName)
	}

	// Open the state file corresponding to the network
	stateFile, err := os.Open(stateFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open state file %s: %w", stateFilePath, err)
	}
	defer stateFile.Close()
	fileContents, err := ioutil.ReadAll(stateFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read state file: %w", err)
	}

	if len(fileContents) == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	// Decode the network state from the file
	var state models.Network

	if err := json.Unmarshal(fileContents, &state); err != nil {
		return nil, fmt.Errorf("failed to decode network state from file %s: %w", stateFilePath, err)
	}

	// Return the decoded network state
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
