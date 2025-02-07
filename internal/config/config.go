package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Aperocky/git-paca/internal/types"
)

const (
	defaultUrl       = "http://localhost:11434"
	defaultModelName = "qwen2.5-coder:14b-instruct-q6_K"
	defaultMaxCtx    = 4096
)

func getDefaultConfig() *types.PacaConfig {
	return &types.PacaConfig{
		Url:       defaultUrl,
		ModelName: defaultModelName,
		MaxCtx:    32768,
		Options:   getDefaultOllamaOptions(),
	}
}

func getDefaultOllamaOptions() *types.OllamaOptions {
	return &types.OllamaOptions{
		NumCtx:        4096, // Reasonable context window for code snippets
		Temperature:   0.3,  // Lower temperature for more focused/deterministic output
		TopP:          0.9,  // Slightly constrained but still allows some flexibility
		RepeatPenalty: 1.1,  // Slight penalty to discourage repetition
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

func LoadConfig() (*types.PacaConfig, error) {
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

	var config *types.PacaConfig
	if err := json.Unmarshal(data, &config); err != nil {
		// If parsing fails, return default config
		return nil, fmt.Errorf("error parsing config file: %v", err)
	}

	if config.MaxCtx == 0 {
		config.MaxCtx = 4096
	}
	if config.Url == "" {
		return nil, fmt.Errorf("url cannot be empty")
	}
	if config.ModelName == "" {
		return nil, fmt.Errorf("model name cannot be empty")
	}

	// Avoid null pointer when config does not have options.
	if config.Options == nil {
		config.Options = getDefaultOllamaOptions()
	}
	return config, nil
}
