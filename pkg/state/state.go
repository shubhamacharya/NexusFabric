package state

import (
	"encoding/json"
	"os"
)

type NetworkState struct {
	Networks map[string]interface{} `json:"networks"`
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
