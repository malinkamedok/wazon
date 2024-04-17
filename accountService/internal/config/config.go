package config

import (
	"github.com/caarlos0/env/v7"
	"log"
)

type Config struct {
	Port           string `env:"ACCOUNT_PORT" envDefault:"8002"`
	PostgresUrl    string `env:"WAZON_DB_URL"`
	StorefrontUrl  string `env:"STOREFRONT_URL" envDefault:"http://storefront:"`
	StoreFrontPort string `env:"STOREFRONT_PORT" envDefault:"8000"`
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
