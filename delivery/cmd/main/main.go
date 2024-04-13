package main

import (
	"delivery/internal/app"
	"delivery/internal/config"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
