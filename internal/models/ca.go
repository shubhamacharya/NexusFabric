package model

type Ca struct {
	Name                                  string  `yaml:"name" json:"name"`
	FabricCAHome                          string  `yaml:"FABRIC_CA_HOME" json:"FABRIC_CA_HOME"`
	FabricCAServerCAName                  *string `yaml:"FABRIC_CA_SERVER_CA_NAME,omitempty" json:"FABRIC_CA_SERVER_CA_NAME,omitempty"`
	FabricCAServerOperationsListenAddress string  `yaml:"FABRIC_CA_SERVER_OPERATIONS_LISTENADDRESS" json:"FABRIC_CA_SERVER_OPERATIONS_LISTENADDRESS"`
	FabricCAServerPort                    int     `yaml:"FABRIC_CA_SERVER_PORT" json:"FABRIC_CA_SERVER_PORT"`
	FabricCAServerTLSEnabled              bool    `yaml:"FABRIC_CA_SERVER_TLS_ENABLED" json:"FABRIC_CA_SERVER_TLS_ENABLED"`
	Volumes                               *string `yaml:"volumes,omitempty" json:"volumes,omitempty"`
	ServerPort                            int     `yaml:"serverport" json:"serverport"`
	OperationsListenPort                  int     `yaml:"operationslistenport" json:"operationslistenport"`
}

func NewCa() Ca {
	return Ca{
		Name:                                  "",
		FabricCAHome:                          "/etc/hyperledger/fabric-ca-server/crypto",
		FabricCAServerOperationsListenAddress: "0.0.0.0:18054",
		FabricCAServerPort:                    8054,
		FabricCAServerTLSEnabled:              true,
		ServerPort:                            8054,
		OperationsListenPort:                  18054,
	}
}
