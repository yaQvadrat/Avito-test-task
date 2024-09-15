package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"log"`
		PG   `yaml:"postgres"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"`
		Version string `env-required:"true" yaml:"version"`
	}

	HTTP struct {
		Address string `env-required:"true" env:"SERVER_ADDRESS"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"level"`
	}

	PG struct {
		MaxPoolSize int    `env-required:"true" yaml:"max_pool_size"`
		URL         string `env-required:"true" env:"POSTGRES_CONN"`
	}
)

func New(configPath string) (*Config, error) {
	cfg := &Config{}

	// Reading config from .yaml file and enviroment (env variables more important)
	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		return nil, fmt.Errorf("config - New - cleanenv.ReadConfig: %w", err)
	}

	return cfg, nil
}
