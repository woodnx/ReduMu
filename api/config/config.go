package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env        string `env:"REDUMU_ENV" envDefault:"dev"`
	Port       int    `env:"PORT" envDefault:"61106"`
	DBHost     string `env:"DB_HOST" envDefault:"db"`
	DBPort     int    `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER" envDefault:"postgres"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"postgres"`
	DBName     string `env:"DB_NAME" envDefault:"redumu"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
