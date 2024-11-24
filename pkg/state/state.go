package state

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/shubhamacharya/NexusFabric/internal/constants"
	models "github.com/shubhamacharya/NexusFabric/internal/models"
)

type NetworkState struct {
	Networks map[string]interface{} `json:"networks"`
}

func CreateState(networkConfig *models.Config) {
	// fmt.Printf("%v : ", networkConfig.Fabric)
	var network models.Network
	network.Name = networkConfig.Fabric.Netname
	if networkConfig.Fabric.StartPort == 0 {
		network.Port = models.NewPort(constants.DefaultPort)
	} else {
		network.Port = models.NewPort(networkConfig.Fabric.StartPort)
	}
	present := map[string]bool{}
	for _, v := range networkConfig.Fabric.Peers {
		org := strings.Split(v, ".")[1]
		if present[org] != true {
			present[org] = true

			tempOrg := models.BuildNewOrganization(org, "ca-"+org, constants.TLSEnabled,&network.Port,)

			// network.Organizations = append(network.Organizations, org)
			fmt.Println(tempOrg)
		}
	}
}
func LoadState(filePath string) (*NetworkState, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var state NetworkState
	if err := json.NewDecoder(file).Decode(&state); err != nil {
		return nil, err
	}
	return &state, nil
}

func SaveState(filePath string, state *NetworkState) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(state)
}
