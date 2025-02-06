package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type PacaConfig struct {
	Url       string `json:"url"`
	ModelName string `json:"model"`
}

func getDefaultConfig() PacaConfig {
	return PacaConfig{
		Url:       "http://localhost:11434",
		ModelName: "qwen2.5-coder:14b-instruct-q6_K",
	}
}

func createConfig(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	defaultConfig := getDefaultConfig()
	configJson, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, configJson, 0644)
}

func LoadConfig() (PacaConfig, error) {
	homeDir, _ := os.UserHomeDir()
	configPath := filepath.Join(homeDir, ".config", "paca", "paca-config.json")
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			err := createConfig(configPath)
			if err != nil {
				fmt.Printf("Unable to create config file at %v", err)
			}
		}
		return getDefaultConfig(), nil
	}

	var config PacaConfig
	if err := json.Unmarshal(data, &config); err != nil {
		// If parsing fails, return default config
		return getDefaultConfig(), fmt.Errorf("error parsing config file: %v", err)
	}
	return config, nil
}
