package config

import (
	"encoding/json"
	"os"
)

// Reads the json config file for any service and parses
// the data into the config struct for that service
func Load[T any](path string, cfg *T) error {
	jsonBytes, err := os.ReadFile(path)		
	if err != nil {
		return err
	}
	
	return json.Unmarshal(jsonBytes, cfg)
}
