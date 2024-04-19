package config

import (
	"go.uber.org/zap"
	"storefront/pkg/logger"

	"github.com/caarlos0/env/v7"
)

type Config struct {
	Port        string `env:"PORT" envDefault:"8082"`
	PostgresUrl string `env:"POSTGRES_URL"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		logger.Error("error in parsing env", zap.Error(err))
		return nil, err
	}
	return cfg, nil
}
