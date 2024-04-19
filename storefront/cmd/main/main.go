package main

import (
	"go.uber.org/zap"
	"storefront/internal/app"
	"storefront/internal/config"
	"storefront/pkg/logger"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatal("config init error", zap.Error(err))
	}

	app.Run(cfg)
}
