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

	mux.Get("/", app.Home)
	mux.NotFound(app.NotFound)

	// Register Routes
	mux.Route("/register", func(mux chi.Router) {
		mux.Get("/", app.RegisterPage)
		mux.Post("/save", app.SaveNewUser)
	})

	// AUTH Routes
	mux.Route("/login", func(mux chi.Router) {
		mux.Get("/", app.LoginPage)
		mux.Post("/authenticate", app.CreateAuthToken)
	})

	// Plant Routes
	mux.Route("/plant", func(mux chi.Router) {
		mux.Get("/", app.PlantPage)
		mux.Post("/save", app.PlantPage)
		mux.Delete("/delete", app.PlantPage)
		mux.Put("/update", app.PlantPage)
	})

	// Users Routes
	mux.Route("/users", func(mux chi.Router) {
		mux.Get("/", app.UserPage)
	})

	// FHX Routes
	mux.Route("/fhx", func(mux chi.Router) {
		mux.Get("/", app.FhxPage)
	})

	fileServer := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/*", http.StripPrefix("/assets", fileServer))

	return mux
}
