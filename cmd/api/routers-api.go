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

	mux.Post("/api/read-fhx", app.ReadFhx)
	mux.Get("/api/all-plants", app.AllPlants)
	mux.Post("/api/allGetOperations", app.GetOperations)
	mux.Post("/api/getParamsFromOPId", app.getParamsFromOPId)

	return mux
}
