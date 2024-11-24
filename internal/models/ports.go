package model

type Port struct {
	Port int `yaml:"port" json:"port"`
}

func NewPort(initPort int) *Port {
	return &Port{Port: initPort}
}

func GetNextAvailablePort(portObj *Port) {
	portObj.Port++
}
