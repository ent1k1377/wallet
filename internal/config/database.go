package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type DatabaseConfig struct {
	Username string `env:"POSTGRES_USERNAME, required"`
	Password string `env:"POSTGRES_PASSWORD, required"`
	Name     string `env:"POSTGRES_DB, required"`
	Host     string `env:"DB_HOST, required"`
	Port     string `env:"DB_PORT, required"`
}

func LoadDatabaseConfig() (*DatabaseConfig, error) {
	var config DatabaseConfig
	err := env.Parse(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.Username, c.Password, c.Host, c.Port, c.Name)
}
