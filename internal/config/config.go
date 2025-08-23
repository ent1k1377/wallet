package config

type Config struct {
	ServerConfig   *ServerConfig
	DatabaseConfig *DatabaseConfig
}

func LoadConfig() (*Config, error) {
	serverConfig, err := LoadServerConfig()
	if err != nil {
		return nil, err
	}

	databaseConfig, err := LoadDatabaseConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		ServerConfig:   serverConfig,
		DatabaseConfig: databaseConfig,
	}, nil
}

func MustLoadConfig() *Config {
	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	return cfg
}
