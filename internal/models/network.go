package model

type Network struct {
	// DomainName    string         `yaml:"domainName" json:"domainName"`       // Domain name
	CA            Ca             `yaml:"ca" json:"ca"`                       // Certificate Authority (CA) details
	CAOrderer     Ca             `yaml:"caorderer" json:"caorderer"`         // CA for the orderer
	Orderers      []Orderer      `yaml:"orderer" json:"orderer"`             // Orderer details
	QtyOrgs       int            `yaml:"qtyorgs" json:"qtyorgs"`             // Number of organizations
	QtyOrds       int            `yaml:"qtyords" json:"qtyords"`             // Number of orderers
	Organizations []Organization `yaml:"organizations" json:"organizations"` // List of organizations
	Name          string         `yaml:"name" json:"name"`                   // network name
	Chaincodes    []Chaincode    `yaml:"chaincodes" json:"chaincodes"`       // List of chaincodes
	Port          *Port          `yaml:"port" json:"port"`
}
