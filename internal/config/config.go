package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"urlshortener/pkg/postgres"
)

type Config struct {
	Postgres postgres.Config `yaml:"POSTGRES" env:"POSTGRES_HOST" env-default:"localhost"`
	RESTPort string          `yaml:"REST_PORT" env:"REST_PORT" env-default:"8080"`
}

func New() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig("./config/config.yaml", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
