package model

// Ca represents the configuration for a Certificate Authority (CA) in Hyperledger Fabric.
type Ca struct {
	Name                                  string  `yaml:"name" json:"name"`                                                                           // Name of the CA
	FabricCAHome                          string  `yaml:"FABRIC_CA_HOME" json:"FABRIC_CA_HOME"`                                                       // Path to the CA home directory
	FabricCAServerCAName                  *string `yaml:"FABRIC_CA_SERVER_CA_NAME,omitempty" json:"FABRIC_CA_SERVER_CA_NAME,omitempty"`               // Optional CA server name
	FabricCAServerOperationsListenAddress string  `yaml:"FABRIC_CA_SERVER_OPERATIONS_LISTENADDRESS" json:"FABRIC_CA_SERVER_OPERATIONS_LISTENADDRESS"` // Address for operations
	FabricCAServerPort                    int     `yaml:"FABRIC_CA_SERVER_PORT" json:"FABRIC_CA_SERVER_PORT"`                                         // Port for the CA server
	FabricCAServerTLSEnabled              bool    `yaml:"FABRIC_CA_SERVER_TLS_ENABLED" json:"FABRIC_CA_SERVER_TLS_ENABLED"`                           // Whether TLS is enabled
	Volumes                               *string `yaml:"volumes,omitempty" json:"volumes,omitempty"`                                                 // Optional volumes for the CA
	ServerPort                            int     `yaml:"serverport" json:"serverport"`                                                               // Server port for the CA
	OperationsListenPort                  int     `yaml:"operationslistenport" json:"operationslistenport"`                                           // Port for operations
}

// Default constants for CA configuration
const (
	DefaultFabricCAHome                          = "/etc/hyperledger/fabric-ca-server/crypto"
	DefaultFabricCAServerOperationsListenAddress = "0.0.0.0:18054"
	DefaultFabricCAServerPort                    = 8054
	DefaultOperationsListenPort                  = 18054
)

// NewCa initializes a CA instance with default values.
func NewCa() Ca {
	return Ca{
		Name:                                  "", // Default to an empty name
		FabricCAHome:                          DefaultFabricCAHome,
		FabricCAServerOperationsListenAddress: DefaultFabricCAServerOperationsListenAddress,
		FabricCAServerPort:                    DefaultFabricCAServerPort,
		FabricCAServerTLSEnabled:              true,
		ServerPort:                            DefaultFabricCAServerPort,
		OperationsListenPort:                  DefaultOperationsListenPort,
	}
}

// BuildNewCA creates a new CA instance with the given parameters.
// Parameters:
// - Name: The name of the CA.
// - FabricCAServerTLSEnabled: Indicates whether TLS is enabled.
// - serverPort: The server port for the CA.
// - opsListenPort: The port for operations.
// Returns:
// - A configured CA instance.
func BuildNewCA(Name string, FabricCAServerTLSEnabled bool, port *Port) Ca {
	// Validate input ports
	GetNextAvailablePort(port)
	serverPort := port.Port
	GetNextAvailablePort(port)
	opsListenPort := port.Port

	if serverPort <= 0 || opsListenPort <= 0 {
		panic("serverPort and opsListenPort must be positive integers")
	}

	return Ca{
		Name:                                  Name,
		FabricCAHome:                          DefaultFabricCAHome,
		FabricCAServerOperationsListenAddress: DefaultFabricCAServerOperationsListenAddress,
		FabricCAServerTLSEnabled:              FabricCAServerTLSEnabled,
		ServerPort:                            serverPort,
		OperationsListenPort:                  opsListenPort,
	}
}
