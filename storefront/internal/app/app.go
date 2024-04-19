package app

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"storefront/internal/config"
	v1 "storefront/internal/controller/v1"
	"storefront/internal/usecase"
	"storefront/internal/usecase/repo"
	"storefront/pkg/httpserver"
	"storefront/pkg/logger"
	"storefront/pkg/postgres"
	"syscall"
)

func Run(cfg *config.Config) {

	logger.InitLogger()

	pg, err := postgres.New(cfg)
	if err != nil {
		logger.Fatal("postgres connection error", zap.Error(err))
		return
	}

	s := usecase.NewStorefrontUseCase(repo.NewPostgresRepo(pg))

	handler := chi.NewRouter()

	v1.NewRouter(handler, s)

	server := httpserver.New(handler, httpserver.Port(cfg.Port))
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)
	logger.Info("storefront server started", zap.String("port", cfg.Port))

	select {
	case s := <-interruption:
		logger.Warn("Interruption channel: ", zap.String("notification", s.String()))
	case err := <-server.Notify():
		logger.Warn("Server notify channel: ", zap.Error(err))
	}

	err = server.Shutdown()
	if err != nil {
		logger.Error("Error shutting down server: ", zap.Error(err))
	}
}
