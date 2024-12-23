package model

import (
	"path/filepath"
)

type Paths struct {
	APPPATH                  string `yaml:"APPPATH" json:"APPPATH"`
	GENERICNETCONFIG         string `yaml:"GENERICNETCONFIG" json:"GENERICNETCONFIG"`
	NETWORKPATH              string `yaml:"NETWORKPATH" json:"NETWORKPATH"`
	NETWORKSTATEPATH         string `yaml: "NETWORKSTATEPATH" json:"NETWORKSTATEPATH"`
	NETWORKCONFIGPATH        string `yaml:"NETWORKCONFIGPATH" json:"NETWORKCONFIGPATH"`
	NETWORKCONFIGTXFILE      string `yaml:"NETWORKCONFIGTXFILE" json:"NETWORKCONFIGTXFILE"`
	NETWORKCONFIGTXJSONFILE  string `yaml:"NETWORKCONFIGTXJSONFILE" json:"NETWORKCONFIGTXJSONFILE"`
	NETWORKCONFIGBUILDPATH   string `yaml:"NETWORKCONFIGBUILDPATH" json:"NETWORKCONFIGBUILDPATH"`
	FABRICCAPATH             string `yaml:"FABRICCAPATH" json:"FABRICCAPATH"`
	CANETWORKNAME            string `yaml:"CANETWORKNAME" json:"CANETWORKNAME"`
	CANETWORKPATH            string `yaml:"CANETWORKPATH" json:"CANETWORKPATH"`
	CANETWORKCRYPTOPATH      string `yaml:"CANETWORKCRYPTOPATH" json:"CANETWORKCRYPTOPATH"`
	CACLIENTNETWORKPATH      string `yaml:"CACLIENTNETWORKPATH" json:"CACLIENTNETWORKPATH"`
	CACLIENTNETWORKMSPPATH   string `yaml:"CACLIENTNETWORKMSPPATH" json:"CACLIENTNETWORKMSPPATH"`
	CACERTNETWORKFILE        string `yaml:"CACERTNETWORKFILE" json:"CACERTNETWORKFILE"`
	TLSCERTNETWORKFILE       string `yaml:"TLSCERTNETWORKFILE" json:"TLSCERTNETWORKFILE"`
	CAORDERERNAME            string `yaml:"CAORDERERNAME" json:"CAORDERERNAME"`
	CAORDERERPATH            string `yaml:"CAORDERERPATH" json:"CAORDERERPATH"`
	CAORDERERCRYPTOPATH      string `yaml:"CAORDERERCRYPTOPATH" json:"CAORDERERCRYPTOPATH"`
	CAORDERERCACLIENTPATH    string `yaml:"CAORDERERCACLIENTPATH" json:"CAORDERERCACLIENTPATH"`
	CAORDERERCACLIENTMSPPATH string `yaml:"CAORDERERCACLIENTMSPPATH" json:"CAORDERERCACLIENTMSPPATH"`
	CACERTORDERERFILE        string `yaml:"CACERTORDERERFILE" json:"CACERTORDERERFILE"`
	TLSCERTORDERERFILE       string `yaml:"TLSCERTORDERERFILE" json:"TLSCERTORDERERFILE"`
	ORDERERNAME              string `yaml:"ORDERERNAME" json:"ORDERERNAME"`
	ORDERERORGPATH           string `yaml:"ORDERERORGPATH" json:"ORDERERORGPATH"`
	ORDERERORGADMINPATH      string `yaml:"ORDERERORGADMINPATH" json:"ORDERERORGADMINPATH"`
	ORDERERORGMSPPATH        string `yaml:"ORDERERORGMSPPATH" json:"ORDERERORGMSPPATH"`
	ORDERERORGSIGNCERTPATH   string `yaml:"ORDERERORGSIGNCERTPATH" json:"ORDERERORGSIGNCERTPATH"`
	ORDNETWORKPATH           string `yaml:"ORDNETWORKPATH" json:"ORDNETWORKPATH"`
	ORDNETWORKMSPPATH        string `yaml:"ORDNETWORKMSPPATH" json:"ORDNETWORKMSPPATH"`
	ORDTLSCAMSPPATH          string `yaml:"ORDTLSCAMSPPATH" json:"ORDTLSCAMSPPATH"`
	ORDNETWORKCACERTPATH     string `yaml:"ORDNETWORKCACERTPATH" json:"ORDNETWORKCACERTPATH"`
	ORDSIGNCERTMSPPATH       string `yaml:"ORDSIGNCERTMSPPATH" json:"ORDSIGNCERTMSPPATH"`
	ORDKEYSTOREMSPPATH       string `yaml:"ORDKEYSTOREMSPPATH" json:"ORDKEYSTOREMSPPATH"`
	ORDNETWORKTLSPATH        string `yaml:"ORDNETWORKTLSPATH" json:"ORDNETWORKTLSPATH"`
	ORDNETWORKADMINCERTPATH  string `yaml:"ORDNETWORKADMINCERTPATH" json:"ORDNETWORKADMINCERTPATH"`
	ORDTLSCAPATH             string `yaml:"ORDTLSCAPATH" json:"ORDTLSCAPATH"`
	ORDSIGNCERTPATH          string `yaml:"ORDSIGNCERTPATH" json:"ORDSIGNCERTPATH"`
	ORDKEYSTOREPATH          string `yaml:"ORDKEYSTOREPATH" json:"ORDKEYSTOREPATH"`
	PEERORGPATH              string `yaml:"PEERORGPATH" json:"PEERORGPATH"`
	CCCRYPTOPATH             string `yaml:"CCCRYPTOPATH" json:"CCCRYPTOPATH"`
	CCCRYPTOTLSPATH          string `yaml:"CCCRYPTOTLSPATH" json:"CCCRYPTOTLSPATH"`
	CCCRYPTOTLSCAPATH        string `yaml:"CCCRYPTOTLSCAPATH" json:"CCCRYPTOTLSCAPATH"`
	CHANNELARTIFACTSPATH     string `yaml:"CHANNELARTIFACTSPATH" json:"CHANNELARTIFACTSPATH"`
	BLOCKFILE                string `yaml:"BLOCKFILE" json:"BLOCKFILE"`
	COMPOSEPATH              string `yaml:"COMPOSEPATH" json:"COMPOSEPATH"`
	CONFIGPATH               string `yaml:"CONFIGPATH" json:"CONFIGPATH"`
	CONFIGPEER               string `yaml:"CONFIGPEER" json:"CONFIGPEER"`
	CONFIGTX                 string `yaml:"CONFIGTX" json:"CONFIGTX"`
	CONFIGORDERER            string `yaml:"CONFIGORDERER" json:"CONFIGORDERER"`
	CLIEXTPATH               string `yaml:"CLIEXTPATH" json:"CLIEXTPATH"`
	CLIPATH                  string `yaml:"CLIPATH" json:"CLIPATH"`
	EXTCONFIGTX              string `yaml:"EXTCONFIGTX" json:"EXTCONFIGTX"`
	CLIBLOCKPATH             string `yaml:"CLIBLOCKPATH" json:"CLIBLOCKPATH"`
	CLIROOTCA                string `yaml:"CLIROOTCA" json:"CLIROOTCA"`
	CLISERVERCRT             string `yaml:"" json:""`
	CLISERVERKEY             string `yaml:"" json:""`
	PEERCFGPATH              string `yaml:"" json:""`
	ORGPATH                  string `yaml:"" json:""`
	ORGCRYPTOPATH            string `yaml:"" json:""`
	ORGCACLIENTPATH          string `yaml:"" json:""`
	ORGMSPPATH               string `yaml:"" json:""`
	CAORGNAME                string `yaml:"" json:""`
	CAORGPATH                string `yaml:"" json:""`
	CAORGCRYPTOPATH          string `yaml:"" json:""`
	CAORGCACLIENTPATH        string `yaml:"" json:""`
	CAORGCACLIENTMSPPATH     string `yaml:"" json:""`
	TLSCERTORGFILE           string `yaml:"" json:""`
	CACERTORGFILE            string `yaml:"" json:""`
	ORGNAME                  string `yaml:"" json:""`
	ORGADMINCERTPATH         string `yaml:"" json:""`
	PEERPATH                 string `yaml:"" json:""`
	PEERNAME                 string `yaml:"" json:""`
	PEERMSPPATH              string `yaml:"" json:""`
	PEERTLSCAMSPPATH         string `yaml:"" json:""`
	PEERTLSPATH              string `yaml:"" json:""`
	PEERCACERTPATH           string `yaml:"" json:""`
	PEERTLSCAPATH            string `yaml:"" json:""`
	PEERSIGNCERTPATH         string `yaml:"" json:""`
	PEERKEYSTOREPATH         string `yaml:"" json:""`
	PEERADMINCERTPATH        string `yaml:"" json:""`
	CHAINCODEPATH            string `yaml:"" json:""`
	CHAINCODEBUILDPATH       string `yaml:"" json:""`
	CHAINCODESRC             string `yaml:"" json:""`
	CHAINCODEPKG             string `yaml:"" json:""`
	PEERSERVERCRT            string `yaml:"" json:""`
	PEERSERVERKEY            string `yaml:"" json:""`
	PEERCAROOT               string `yaml:"" json:""`
	ORDERERORGADMINCERTPATH  string `yaml:"" json:""`
	ORDERERORGTLSCAMSPPATH   string `yaml:"" json:""`
	ORGTLSPATH               string `yaml:"ORGTLSPATH" json:"ORGTLSPATH"`
	ORGTLSTLSCAPATH          string `yaml:"ORGTLSTLSCAPATH" json:"ORGTLSTLSCAPATH"`
	ORGSIGNCERTPATH          string `yaml:"ORGSIGNCERTPATH" json:"ORGSIGNCERTPATH"`
	ORDERERORGCAMSPPATH      string `yaml:"ORDERERORGCAMSPPATH" json:"ORDERERORGCAMSPPATH"`
	CLIHOSTNAME              string `yaml:"CLIHOSTNAME" json:"CLIHOSTNAME"`
}

func NewPath(network *Network) Paths {
	Path := Paths{}
	// Application Path : ${PWD}
	Path.APPPATH, _ = filepath.Abs(".")
	// ${PWD}/network_config.yaml
	Path.GENERICNETCONFIG = Path.APPPATH + "/network_config.yaml"
	// NETWORK Path : ${PWD}/[NETWORKNAME]/
	Path.NETWORKPATH = Path.APPPATH + "/networks/" + network.Name + "/"
	// Network State file
	Path.NETWORKSTATEPATH = Path.NETWORKPATH + "state.json"

	// ${PWD}/networks/[NETWORKNAME]/config/
	Path.NETWORKCONFIGPATH = Path.NETWORKPATH + "config/"
	// ${PWD}/networks/[NETWORKNAME]/config/configtx.yaml
	Path.NETWORKCONFIGTXFILE = Path.NETWORKCONFIGPATH + "configtx.yaml"
	// ${PWD}/networks/[NETWORKNAME]/config/configtx.json
	Path.NETWORKCONFIGTXJSONFILE = Path.NETWORKCONFIGPATH + "configtx.json"

	// ${PWD}/networks/[NETWORKNAME]/config/build/
	Path.NETWORKCONFIGBUILDPATH = Path.NETWORKCONFIGPATH + "build/"

	// Fabric CAs Paths:  ${PWD}/networks/[NETWORKNAME]/fabricca/
	Path.FABRICCAPATH = Path.NETWORKPATH + "fabric-ca/"

	// Fabric CA Orderer : ca.orderer.[DOMAIN]
	Path.CAORDERERNAME = network.CAOrderer.Name
	// ${PWD}/networks/[NETWORK]/fabric-ca/ca.orderer.[DOMAIN]/
	Path.CAORDERERPATH = Path.FABRICCAPATH + network.CAOrderer.Name + "/"
	// ${PWD}/networks/[NETWORK]/fabric-ca/ca.orderer.[DOMAIN]/admin/
	Path.CAORDERERCACLIENTPATH = Path.CAORDERERPATH + "admin/"
	// ${PWD}/networks/[NETWORK]/fabric-ca/ca.orderer.[DOMAIN]/admin/msp/
	Path.CAORDERERCACLIENTMSPPATH = Path.CAORDERERCACLIENTPATH + "msp/"

	// Fabric CA Domain Cert Files
	// ${PWD}/networks/[NETWORK]/fabric-ca/ca.orderer.[DOMAIN]/ca-cert.pem
	Path.CACERTORDERERFILE = Path.CAORDERERCRYPTOPATH + "ca-cert.pem"
	// ${PWD}/networks/[NETWORK]/fabricca/ca.orderer.[DOMAIN]/tls-cert.pem
	Path.TLSCERTORDERERFILE = Path.CAORDERERCRYPTOPATH + "tls-cert.pem"

	return Path
}
