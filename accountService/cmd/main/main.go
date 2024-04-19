package main

import (
	"accountservice/internal/app"
	"accountservice/internal/config"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
