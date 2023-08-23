package config

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"go-wp/pkg/filesystem"
	"io"
	"os"
	"path"
	"path/filepath"
)

func Get() *AppConfig {
	path, _ := filepath.Abs("config/config.json")
	fmt.Println(path)
	jsonFile, _ := os.Open(path)

	bytes, _ := io.ReadAll(jsonFile)

	var config AppConfig

	json.Unmarshal(bytes, &config)
	fmt.Printf("Using configuration: %v\n", config)
	return &config
}

var (
	//go:embed config.json
	configFile []byte
)

func EnsureConfig() bool {

	filename := "config/config.json"
	filepath := path.Dir(filename)

	if filesystem.Exists(filename) {
		return true
	}

	if !filesystem.ExistsDir(filepath) {
		err := os.Mkdir(filepath, 0644)
		if err != nil {
			fmt.Printf("Could not create directory %s, %s\n", filepath, err)
			os.Exit(1)
		}
	}

	err := os.WriteFile(filename, configFile, 0644)
	if err != nil {
		fmt.Printf("Could not generate config file %s, %s\n", filename, err)
		os.Exit(1)
	}
	return false
}
