package model

import "strconv"

type Orderer struct {
	Name                               string   `yaml:"name" json:"name"`
	FabricCfgPath                      string   `yaml:"FABRIC_CFG_PATH" json:"FABRIC_CFG_PATH"`
	FabricLoggingSpec                  string   `yaml:"FABRIC_LOGGING_SPEC" json:"FABRIC_LOGGING_SPEC"`
	GeneralListenPort                  int      `yaml:"ORDERER_GENERAL_LISTENPORT" json:"ORDERER_GENERAL_LISTENPORT"`
	OrdererGeneralListenAddress        string   `yaml:"ORDERER_GENERAL_LISTENADDRESS" json:"ORDERER_GENERAL_LISTENADDRESS"`
	OrdererGeneralLocalMSPID           string   `yaml:"ORDERER_GENERAL_LOCALMSPID" json:"ORDERER_GENERAL_LOCALMSPID"`
	OrdererGeneralLocalMSPDir          string   `yaml:"ORDERER_GENERAL_LOCALMSPDIR" json:"ORDERER_GENERAL_LOCALMSPDIR"`
	OrdererGeneralTLSEnabled           bool     `yaml:"ORDERER_GENERAL_TLS_ENABLED" json:"ORDERER_GENERAL_TLS_ENABLED"`
	OrdererGeneralTLSCertificate       string   `yaml:"ORDERER_GENERAL_TLS_CERTIFICATE" json:"ORDERER_GENERAL_TLS_CERTIFICATE"`
	OrdererGeneralTLSPrivateKey        string   `yaml:"ORDERER_GENERAL_TLS_PRIVATEKEY" json:"ORDERER_GENERAL_TLS_PRIVATEKEY"`
	OrdererGeneralTLSRootCAs           string   `yaml:"ORDERER_GENERAL_TLS_ROOTCAS" json:"ORDERER_GENERAL_TLS_ROOTCAS"`
	OrdererChannelParticipationEnabled bool     `yaml:"ORDERER_CHANNELPARTICIPATION_ENABLED" json:"ORDERER_CHANNELPARTICIPATION_ENABLED"`
	OperationsListenPort               int      `yaml:"operationslistenport" json:"operationslistenport"`
	OrdererOperationsListenAddress     string   `yaml:"ORDERER_OPERATIONS_LISTENADDRESS,omitempty" json:"ORDERER_OPERATIONS_LISTENADDRESS,omitempty"`
	AdminListenPort                    int      `yaml:"adminlistenport" json:"adminlistenport"`
	OrdererAdminListenAddress          string   `yaml:"ORDERER_ADMIN_LISTENADDRESS" json:"ORDERER_ADMIN_LISTENADDRESS"`
	OrdererAdminTLSEnabled             bool     `yaml:"ORDERER_ADMIN_TLS_ENABLED" json:"ORDERER_ADMIN_TLS_ENABLED"`
	OrdererAdminTLSCertificate         string   `yaml:"ORDERER_ADMIN_TLS_CERTIFICATE" json:"ORDERER_ADMIN_TLS_CERTIFICATE"`
	OrdererAdminTLSClientRootCAs       string   `yaml:"ORDERER_ADMIN_TLS_CLIENTROOTCAS" json:"ORDERER_ADMIN_TLS_CLIENTROOTCAS"`
	OrdererAdminTLSPrivateKey          string   `yaml:"ORDERER_ADMIN_TLS_PRIVATEKEY" json:"ORDERER_ADMIN_TLS_PRIVATEKEY"`
	OrdererAdminTLSRootCAs             string   `yaml:"ORDERER_ADMIN_TLS_ROOTCAS" json:"ORDERER_ADMIN_TLS_ROOTCAS"`
	Volumes                            []string `yaml:"volumes" json:"volumes"`
}

func NewOrderer() Orderer {
	return Orderer{
		FabricCfgPath:                      "/var/hyperledger/fabric/config",
		FabricLoggingSpec:                  "WARN:cauthdsl=debug:policies=debug:msp=debug",
		OrdererGeneralListenAddress:        "0.0.0.0",
		OrdererGeneralLocalMSPID:           "OrdererMSP",
		OrdererGeneralLocalMSPDir:          "/var/hyperledger/orderer/msp",
		OrdererGeneralTLSEnabled:           false,
		OrdererGeneralTLSCertificate:       "/var/hyperledger/orderer/tls/signcerts/cert.crt",
		OrdererGeneralTLSPrivateKey:        "/var/hyperledger/orderer/tls/keystore/key.pem",
		OrdererGeneralTLSRootCAs:           "[/var/hyperledger/orderer/tls/tlscacerts/tls-cert.pem]",
		OrdererChannelParticipationEnabled: true,
		OrdererAdminListenAddress:          "0.0.0.0:7050",
		OrdererAdminTLSEnabled:             false,
		OrdererAdminTLSCertificate:         "/var/hyperledger/orderer/tls/signcerts/cert.crt",
		OrdererAdminTLSClientRootCAs:       "[/var/hyperledger/orderer/tls/tlscacerts/tls-cert.pem]",
		OrdererAdminTLSPrivateKey:          "/var/hyperledger/orderer/tls/keystore/key.pem",
		OrdererAdminTLSRootCAs:             "[/var/hyperledger/orderer/tls/tlscacerts/tls-cert.pem]",
		Volumes:                            []string{},
	}
}

func BuildNewOrderer(ordName string, TLSEnabled bool, port *Port) Orderer {

	orderer := NewOrderer()

	orderer.Name = ordName
	GetNextAvailablePort(port)
	orderer.GeneralListenPort = port.Port
	orderer.OrdererGeneralListenAddress = "0.0.0.0:" + strconv.Itoa(orderer.GeneralListenPort)
	orderer.OrdererGeneralTLSEnabled = TLSEnabled
	orderer.OrdererGeneralLocalMSPID = "OrdererMSP"
	GetNextAvailablePort(port)
	orderer.OperationsListenPort = port.Port
	orderer.OrdererOperationsListenAddress = "0.0.0.0:" + strconv.Itoa(orderer.OperationsListenPort)
	GetNextAvailablePort(port)
	orderer.AdminListenPort = port.Port
	orderer.OrdererAdminListenAddress = "0.0.0.0:" + strconv.Itoa(orderer.AdminListenPort)

	return orderer
}
