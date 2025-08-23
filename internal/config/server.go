package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type ServerConfig struct {
	Port string `yaml:"port, required"`
}

func LoadServerConfig() (*ServerConfig, error) {
	var config ServerConfig
	configPath := "configs/server_config.yaml"
	if path := os.Getenv("SERVER_CONFIG_PATH"); path != "" {
		configPath = path
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
