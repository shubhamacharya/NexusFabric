package model

type Organization struct {
	Name     string  `yaml:"name" json:"name"`
	CA       *Ca     `yaml:"ca,omitempty" json:"ca,omitempty"`
	Peers    []Peer  `yaml:"peers" json:"peers"`
	QtyPeers int     `yaml:"qtypeers" json:"qtypeers"`
	KeyStore *string `yaml:"keystore,omitempty" json:"keystore,omitempty"`
}

func NewOrganization() Organization {
	return Organization{
		Peers:    []Peer{}, // Initialize empty slice
		QtyPeers: 0,        // Default to 0
	}
}
