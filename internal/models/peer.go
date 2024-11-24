package model

import (
	"strconv"
)

type Peer struct {
	Name                              string   `yaml:"name" json:"name"`
	CorePeerID                        *string  `yaml:"CORE_PEER_ID,omitempty" json:"CORE_PEER_ID,omitempty"`
	CorePeerAddress                   *string  `yaml:"CORE_PEER_ADDRESS,omitempty" json:"CORE_PEER_ADDRESS,omitempty"`
	CorePeerLocalMSPID                *string  `yaml:"CORE_PEER_LOCALMSPID,omitempty" json:"CORE_PEER_LOCALMSPID,omitempty"`
	CorePeerMSPConfigPath             string   `yaml:"CORE_PEER_MSPCONFIGPATH" json:"CORE_PEER_MSPCONFIGPATH"`
	CoreVMEndpoint                    string   `yaml:"CORE_VM_ENDPOINT" json:"CORE_VM_ENDPOINT"`
	FabricLoggingSpec                 string   `yaml:"FABRIC_LOGGING_SPEC" json:"FABRIC_LOGGING_SPEC"`
	CorePeerTLSEnabled                bool     `yaml:"CORE_PEER_TLS_ENABLED" json:"CORE_PEER_TLS_ENABLED"`
	CorePeerTLSCertFile               string   `yaml:"CORE_PEER_TLS_CERT_FILE" json:"CORE_PEER_TLS_CERT_FILE"`
	CorePeerTLSKeyFile                string   `yaml:"CORE_PEER_TLS_KEY_FILE" json:"CORE_PEER_TLS_KEY_FILE"`
	CorePeerTLSRootCertFile           string   `yaml:"CORE_PEER_TLS_ROOTCERT_FILE" json:"CORE_PEER_TLS_ROOTCERT_FILE"`
	CorePeerGossipUseLeaderElection   bool     `yaml:"CORE_PEER_GOSSIP_USELEADERELECTION" json:"CORE_PEER_GOSSIP_USELEADERELECTION"`
	CorePeerGossipOrgLeader           bool     `yaml:"CORE_PEER_GOSSIP_ORGLEADER" json:"CORE_PEER_GOSSIP_ORGLEADER"`
	CorePeerGossipExternalEndpoint    *string  `yaml:"CORE_PEER_GOSSIP_EXTERNALENDPOINT,omitempty" json:"CORE_PEER_GOSSIP_EXTERNALENDPOINT,omitempty"`
	CorePeerGossipSkipHandshake       bool     `yaml:"CORE_PEER_GOSSIP_SKIPHANDSHAKE" json:"CORE_PEER_GOSSIP_SKIPHANDSHAKE"`
	CorePeerGossipBootstrap           *string  `yaml:"CORE_PEER_GOSSIP_BOOTSTRAP,omitempty" json:"CORE_PEER_GOSSIP_BOOTSTRAP,omitempty"`
	ChaincodeAsAServiceBuilderConfig  string   `yaml:"CHAINCODE_AS_A_SERVICE_BUILDER_CONFIG" json:"CHAINCODE_AS_A_SERVICE_BUILDER_CONFIG"`
	CoreChaincodeExecuteTimeout       string   `yaml:"CORE_CHAINCODE_EXECUTETIMEOUT" json:"CORE_CHAINCODE_EXECUTETIMEOUT"`
	CoreLedgerStateCouchDBConfigAddr  *string  `yaml:"CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS,omitempty" json:"CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS,omitempty"`
	CoreLedgerStateCouchDBConfigPass  string   `yaml:"CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD" json:"CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD"`
	CoreLedgerStateCouchDBConfigUser  string   `yaml:"CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME" json:"CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME"`
	CoreLedgerStateStateDatabase      string   `yaml:"CORE_LEDGER_STATE_STATEDATABASE" json:"CORE_LEDGER_STATE_STATEDATABASE"`
	CoreOperationsListenAddress       *string  `yaml:"CORE_OPERATIONS_LISTENADDRESS,omitempty" json:"CORE_OPERATIONS_LISTENADDRESS,omitempty"`
	CorePeerChaincodeAddress          *string  `yaml:"CORE_PEER_CHAINCODEADDRESS,omitempty" json:"CORE_PEER_CHAINCODEADDRESS,omitempty"`
	CorePeerChaincodeListenAddress    *string  `yaml:"CORE_PEER_CHAINCODELISTENADDRESS,omitempty" json:"CORE_PEER_CHAINCODELISTENADDRESS,omitempty"`
	CorePeerListenAddress             string   `yaml:"CORE_PEER_LISTENADDRESS" json:"CORE_PEER_LISTENADDRESS"`
	CorePeerProfileEnabled            bool     `yaml:"CORE_PEER_PROFILE_ENABLED" json:"CORE_PEER_PROFILE_ENABLED"`
	CoreVMDockerHostConfigNetworkMode *string  `yaml:"CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE,omitempty" json:"CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE,omitempty"`
	FabricCfgPath                     string   `yaml:"FABRIC_CFG_PATH" json:"FABRIC_CFG_PATH"`
	OperationsListenPort              int      `yaml:"operationslistenport" json:"operationslistenport"`
	ChaincodeListenPort               int      `yaml:"chaincodelistenport" json:"chaincodelistenport"`
	PeerListenPort                    int      `yaml:"peerlistenport" json:"peerlistenport"`
	Volumes                           []string `yaml:"volumes" json:"volumes"`
}

func NewPeer() Peer {
	return Peer{
		CorePeerMSPConfigPath:            "/etc/hyperledger/fabric/msp",
		CoreVMEndpoint:                   "unix:///host/var/run/docker.sock",
		FabricLoggingSpec:                "WARN:grpc=debug",
		CorePeerTLSEnabled:               true,
		CorePeerTLSCertFile:              "/etc/hyperledger/fabric/tls/signcerts/cert.crt",
		CorePeerTLSKeyFile:               "/etc/hyperledger/fabric/tls/keystore/key.pem",
		CorePeerTLSRootCertFile:          "/etc/hyperledger/fabric/tls/tlscacerts/tls-cert.pem",
		CorePeerGossipUseLeaderElection:  true,
		CorePeerGossipOrgLeader:          false,
		CorePeerGossipSkipHandshake:      true,
		ChaincodeAsAServiceBuilderConfig: `{"peername":"peername"}`,
		CoreChaincodeExecuteTimeout:      "300s",
		CoreLedgerStateCouchDBConfigPass: "adminpw",
		CoreLedgerStateCouchDBConfigUser: "admin",
		CoreLedgerStateStateDatabase:     "CouchDB",
		CorePeerListenAddress:            "0.0.0.0:7051",
		CorePeerProfileEnabled:           false,
		FabricCfgPath:                    "/etc/hyperledger/peercfg",
		Volumes:                          []string{},
	}
}

func BuildNewPeers(FabricLoggingSpec string, TLSEnabled bool, port *Port, qtyPeers int) []Peer {
	newPeers := make([]Peer, 0, qtyPeers)
	for i := 0; i < qtyPeers; i++ {
		peer := NewPeer()
		GetNextAvailablePort(port)

		peer.FabricLoggingSpec = FabricLoggingSpec
		peer.CorePeerTLSEnabled = TLSEnabled
		peer.CorePeerListenAddress = "0.0.0.0:" + strconv.Itoa(port.Port)
		newPeers = append(newPeers, peer)
	}
	return newPeers
}
