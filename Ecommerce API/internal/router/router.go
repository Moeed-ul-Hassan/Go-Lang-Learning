package router

import (
	"github.com/Moeed-ul-Hassan/EcomAPI/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", handler.HealthCheck)
	})

	return r
}
