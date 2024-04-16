package config

import (
	"github.com/caarlos0/env/v7"
	"log"
)

type Config struct {
	Port          string `env:"PORT" envDefault:"8000"`
	PostgresUrl   string `env:"POSTGRES_URL"`
	StorefrontUrl string `env:"STOREFRONT_URL" envDefault:"http://localhost:8082" `
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
