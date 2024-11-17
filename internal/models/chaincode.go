package model

type Chaincode struct {
	Name        string  `yaml:"name" json:"name"`                                   // Chaincode name
	ServiceName *string `yaml:"servicename,omitempty" json:"servicename,omitempty"` // Optional service name
	Version     int     `yaml:"version" json:"version"`                             // Version number
	PackageID   *string `yaml:"packageid,omitempty" json:"packageid,omitempty"`     // Optional package ID
	CCPort      int     `yaml:"ccport" json:"ccport"`                               // Chaincode port
	Invoke      bool    `yaml:"invoke" json:"invoke"`                               // Whether to invoke the chaincode
	Function    *string `yaml:"function,omitempty" json:"function,omitempty"`       // Optional function name to invoke
	UseTLS      bool    `yaml:"usetls" json:"usetls"`                               // TLS usage flag
	ClientKey   *string `yaml:"client_key,omitempty" json:"client_key,omitempty"`   // Optional client key
	ClientCert  *string `yaml:"client_cert,omitempty" json:"client_cert,omitempty"` // Optional client certificate
	RootCert    *string `yaml:"root_cert,omitempty" json:"root_cert,omitempty"`     // Optional root certificate
}

func NewChaincode() Chaincode {
	return Chaincode{
		Name:        "",
		Version:     0,
		CCPort:      0,
		Invoke:      false,
		UseTLS:      false,
		ServiceName: nil,
		PackageID:   nil,
		Function:    nil,
		ClientKey:   nil,
		ClientCert:  nil,
		RootCert:    nil,
	}
}
