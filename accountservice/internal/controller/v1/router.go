package v1

import (
	"accountservice/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter(handler *chi.Mux, s usecase.AccountServiceContract, rest usecase.IntegrationContract) {
	handler.Route("/accountservice", func(router chi.Router) {
		router.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Access-Control-Allow-Origin", "X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "origin", "x-requested-with"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}))
		NewUserRoutes(router, s, rest)
	})
}
