package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Get() *AppConfig {
	path, _ := filepath.Abs("config.json")
	fmt.Println(path)
	jsonFile, _ := os.Open(path)

	bytes, _ := io.ReadAll(jsonFile)

	var config AppConfig

	json.Unmarshal(bytes, &config)
	fmt.Printf("Using configuration: %v\n", config)
	return &config
}
