package app

import (
	"delivery/internal/config"
	"delivery/internal/controller/mqProducer"
	v1 "delivery/internal/controller/v1"
	"delivery/internal/usecase"
	"delivery/internal/usecase/mq"
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

	var mqConfig config.MQConfig
	mqConfig, err = mqProducer.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer mqConfig.Connection.Close()
	defer mqConfig.Channel.Close()

	s := usecase.NewDeliveryUseCase(repo.NewPostgresRepo(pg))
	m := usecase.NewDeliveryMQUseCase(mq.NewRabbitMQProducer(mqConfig.Channel, mqConfig.Queue))

	handler := chi.NewRouter()

	v1.NewRouter(handler, s, m)

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
		log.Printf("Http server shutdown error: %s\n", err)
	}
}
