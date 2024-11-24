package model

type Config struct {
	Fabric FabricConfig `yaml:"fabric"`
}

type FabricConfig struct {
	CAs       []string `yaml:"cas"`
	Peers     []string `yaml:"peers"`
	Orderers  []string `yaml:"orderers"`
	Settings  Settings `yaml:"settings"`
	Netname   string   `yaml:"netname"`
	StartPort int      `yaml: startport`
	Port      *Port    `yaml:"port" json:"port"`
}

type Settings struct {
	CA      Logging `yaml:"ca"`
	Peer    Logging `yaml:"peer"`
	Orderer Logging `yaml:"orderer"`
}

type Logging struct {
	FabricLoggingSpec string `yaml:"FABRIC_LOGGING_SPEC"`
}
