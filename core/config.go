package core

import (
	"encoding/json"
	"os"
)

// Initialize service configuration via JSON.
// The "source" params refer the directory of the JSON file.
// Ex: ./config/myconfig.json
func NewJSONConfig(source string) *Base {
	file, err := os.ReadFile(source)

	if err != nil {
		return nil
	}

	var base Base

	json.Unmarshal(file, &base)

	return &base
}
