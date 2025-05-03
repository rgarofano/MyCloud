package utils

import (
	"encoding/json"
	"os"
)

// Parses json file into a struct or slice
func LoadJson[T any](path string, cfg *T) error {
	jsonBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonBytes, cfg)
}
