package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Register Routes
	mux.Route("/fhx/", func(mux chi.Router) {
		mux.Get("/upload", app.ReadFhx)
	})
	mux.Route("/plant/", func(mux chi.Router) {
		mux.Get("/all", app.AllPlants)
	})

	// mux.NotFound(app.NotFound)

	return mux
}
