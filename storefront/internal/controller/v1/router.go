package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"storefront/internal/usecase"
)

func NewRouter(handler *chi.Mux, s usecase.StorefrontContract) {
	handler.Route("/storefront", func(r chi.Router) {
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Access-Control-Allow-Origin", "X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "origin", "x-requested-with"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}))
		NewUserRoutes(r, s)
	})
}
