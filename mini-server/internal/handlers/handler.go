package handlers

import (
	"github.com/go-chi/chi/v5"
	chiMiddle "github.com/go-chi/chi/v5/middleware"

	"github.com/BunnyEncoder20/go-lang-study-project/mini-server/internal/middlewares"
)

// NOTE: The Handler function start with a capital letter, making it a Pulic function
// As a function name starting with a small letter means it's private to the package only
func Handler(r *chi.Mux) {
	// Global Middlewares
	r.Use(chiMiddle.StripSlashes)

	// Calling our route:
	r.Route("/account", func(router chi.Router) {
		// Route specific middlewares: Auth Middleware
		router.Use(middlewares.AuthMiddleware)

		// Service/Business logic would be here:
		router.Get("/coins", GetCoinsBalance)
	})
}
