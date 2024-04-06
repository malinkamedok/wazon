package main

import (
	"log"
	"storefront/internal/app"
	"storefront/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
