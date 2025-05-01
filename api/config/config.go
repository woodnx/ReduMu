package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env  string `env:"REDUMU_ENV" envDefault:"dev"`
	Port int    `env:"PORT" envDefault:"61106"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
