package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{app.config.CorsAllowedOrigin},
		AllowedHeaders: []string{"*"},
	}))

	mux.NotFound(app.notFound)
	mux.MethodNotAllowed(app.methodNotAllowed)

	mux.Use(app.recoverPanic)

	mux.Get("/status", app.status)

	mux.Route("/api", func(r chi.Router) {
		mux.Post("/itinerary", app.itineraryHandler)
	})

	return mux
}
