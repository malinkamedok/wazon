package config

import (
	"log"

	"github.com/caarlos0/env/v7"
)

type Config struct {
	Port        string `env:"PORT" envDefault:"8001"`
	PostgresUrl string `env:"POSTGRES_URL"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		log.Println("Error in parsing env")
		return nil, err
	}
	return cfg, nil
}
