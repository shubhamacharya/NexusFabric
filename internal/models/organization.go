package model

import (
	"github.com/shubhamacharya/NexusFabric/internal/constants"
)

type Organization struct {
	Name     string `yaml:"name" json:"name"`
	CA       Ca     `yaml:"ca,omitempty" json:"ca,omitempty"`
	Peers    []Peer `yaml:"peers" json:"peers"`
	QtyPeers int    `yaml:"qtypeers" json:"qtypeers"`
}

func NewOrganization() Organization {
	return Organization{
		Peers:    []Peer{},                  // Initialize empty slice
		QtyPeers: constants.DefaultQtyPeers, // Default quantity of peers
	}
}

func BuildNewOrganization(
	Name string,
	CAName string,
	TLSEnabled bool,
	port *Port,
	QtyPeers int,
) Organization {
	// Validate organization name
	if Name == "" {
		panic("Organization name cannot be empty")
	}

	// Validate CA name
	if CAName == "" {
		panic("CA name cannot be empty")
	}

	// Validate the number of peers
	if QtyPeers <= 0 {
		QtyPeers = constants.DefaultQtyPeers
	}

	return Organization{
		Name:     Name,
		CA:       BuildNewCA(CAName, TLSEnabled, port),
		Peers:    BuildNewPeers(constants.DefaultFabricLoggingSpec, TLSEnabled, port, QtyPeers),
		QtyPeers: QtyPeers,
	}
}
