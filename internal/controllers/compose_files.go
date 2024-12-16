package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	models "github.com/shubhamacharya/NexusFabric/internal/models"
	state "github.com/shubhamacharya/NexusFabric/pkg/state"
	"gopkg.in/yaml.v3"
)

type DockerCompose struct {
	Version  string              `yaml:"version"`
	Networks map[string]Network  `yaml:"networks"`
	Services map[string]Services `yaml:"services"`
}

type Network struct {
	Name string `yaml:"name"`
}

type Services struct {
	Image       string            `yaml:"image"`
	Labels      map[string]string `yaml:"labels"`
	Environment []string          `yaml:"environment"`
	Ports       []string          `yaml:"ports"`
	Command     string            `yaml:"command"`
	Volumes     []string          `yaml:"volumes"`
	Container   string            `yaml:"container_name"`
	Networks    []string          `yaml:"networks"`
}

// BuildNewCAComposeFile creates a new Docker Compose file for the CA service
func BuildNewCAComposeFile(network *models.Network, CAInst *models.Ca) (*DockerCompose, error) {
	if network == nil || CAInst == nil {
		return nil, fmt.Errorf("network or CA instance is nil")
	}

	// Initialize a new DockerCompose struct
	dockerCompose := &DockerCompose{
		Version: "3.7",
		Networks: map[string]Network{
			network.Name: {
				Name: network.Name,
			},
		},
		Services: map[string]Services{},
	}

	// Define common environment variables for the CA service
	environment := []string{
		"FABRIC_CA_HOME=/etc/hyperledger/fabric-ca-server",
		"FABRIC_CA_SERVER_CA_NAME=" + CAInst.Name,
		"FABRIC_CA_SERVER_TLS_ENABLED=" + strconv.FormatBool(CAInst.FabricCAServerTLSEnabled),
		"FABRIC_CA_SERVER_PORT=" + strconv.Itoa(CAInst.ServerPort),
		"FABRIC_CA_SERVER_OPERATIONS_LISTENADDRESS=0.0.0.0:" + strconv.Itoa(CAInst.OperationsListenPort),
	}

	// Define the service details for the CA
	service := Services{
		Image:       "hyperledger/fabric-ca:latest",
		Labels:      map[string]string{"service": "hyperledger-fabric"},
		Environment: environment,
		Ports: []string{
			strconv.Itoa(CAInst.ServerPort) + ":7054",
			strconv.Itoa(CAInst.OperationsListenPort) + ":17054",
		},
		Command: "sh -c 'fabric-ca-server start -b admin:adminpw -d'",
		Volumes: []string{
			fmt.Sprintf("../organizations/fabric-ca/%s:/etc/hyperledger/fabric-ca-server", CAInst.Name),
		},
		Container: CAInst.Name,
		Networks:  []string{network.Name},
	}

	// Add the service to the DockerCompose struct
	dockerCompose.Services[CAInst.Name] = service
	return dockerCompose, nil
}

// BuildAllNetworkComposeFiles builds and saves compose files for all networks
func BuildAllNetworkComposeFiles(netname string) {
	network, err := state.GetNetworkStateByName(netname)
	if err != nil {
		// Log error if fetching network state fails
		log.Printf("Error retrieving network %s: %v", netname, err)
		return
	}

	if network == nil {
		log.Printf("Network %s not found", netname)
		return
	}

	// Generate Docker Compose file for the network's CA
	composeFile, err := BuildNewCAComposeFile(network, &network.CA)
	if err != nil {
		log.Printf("Error building Docker Compose file for network %s: %v", netname, err)
		return
	}

	// Initialize paths and create directories
	networkPaths := models.NewPath(network)

	// Save the compose file to a file (or print it, depending on your need)
	err = SaveDockerComposeToFile(networkPaths.NETWORKCACOMPOSEPATH, composeFile)
	if err != nil {
		log.Printf("Error saving Docker Compose file for network %s: %v", netname, err)
		return
	}

	// Optionally print or log the generated compose file
	fmt.Printf("Docker Compose file generated and saved for network %s: %s\n", netname, networkPaths.COMPOSEPATH)
}

// SaveDockerComposeToFile saves the Docker Compose object to a file
func SaveDockerComposeToFile(filePath string, dockerCompose *DockerCompose) error {
	// You can use a YAML encoder or JSON encoder based on your preference
	// Here, for simplicity, let's assume you are saving it as JSON
	data, err := yaml.Marshal(dockerCompose)
	if err != nil {
		return fmt.Errorf("failed to marshal Docker Compose data: %w", err)
	}

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write Docker Compose file: %w", err)
	}

	return nil
}
