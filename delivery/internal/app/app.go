package app

import (
	"delivery/internal/config"
	v1 "delivery/internal/controller/v1"
	"delivery/internal/usecase"
	"delivery/internal/usecase/repo"
	"delivery/pkg/httpserver"
	"delivery/pkg/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
)

func Run(cfg *config.Config) {
	pg, err := postgres.New(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}

	s := usecase.NewDeliveryUseCase(repo.NewPostgresRepo(pg))

	handler := chi.NewRouter()

	v1.NewRouter(handler, s)

	server := httpserver.New(handler, httpserver.Port(cfg.Port))
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)
	log.Printf("server started")

	select {
	case s := <-interruption:
		log.Printf("signal: " + s.String())
	case err := <-server.Notify():
		log.Printf("Notify from http server: %s\n", err)
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("Http server shutdown")
	}
}
