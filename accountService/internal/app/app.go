package app

import (
	"accountservice/internal/config"
	"accountservice/internal/controller/mq"
	v1 "accountservice/internal/controller/v1"
	"accountservice/internal/usecase"
	"accountservice/internal/usecase/repo"
	storefrontrest "accountservice/internal/usecase/storefrontRest"
	"accountservice/pkg/httpserver"
	"accountservice/pkg/postgres"
	"github.com/go-chi/chi/v5"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {

	pg, err := postgres.New(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}

	s := usecase.NewStorefrontUseCase(repo.NewPostgresRepo(pg))

	r := usecase.NewIntegrationUsecase(storefrontrest.NewStorefrontRest(cfg))

	handler := chi.NewRouter()

	v1.NewRouter(handler, s, r)

	server := httpserver.New(handler, httpserver.Port(cfg.Port))
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)
	log.Printf("server started")

	mqConfig, err := mq.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	msgs, err := mqConfig.Channel.Consume(
		mqConfig.Queue.Name, // queue
		"",                  // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	if err != nil {
		log.Panicf("Failed to register a consumer: %s", err)
		return
	}

	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	<-forever

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
