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

	//mux.Use(middleware.Logger)
	fileServer := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/*", http.StripPrefix("/assets", fileServer))

	mux.Get("/", app.Home)

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
	mux.Route("/plants", func(mux chi.Router) {
		mux.Get("/", app.PlantPage)
		mux.Get("/allPlants", app.AllPlants)
		mux.Post("/save", app.PlantSave)
		mux.Delete("/delete", app.PlantDelete)
		mux.Put("/update", app.PlantUpdate)
	})

	// Users Routes
	mux.Route("/users", func(mux chi.Router) {
		mux.Get("/", app.UserPage)
	})

	// FHX Routes
	mux.Route("/fhx", func(mux chi.Router) {
		mux.Use(app.Plant)
		mux.Get("/", app.FhxPage)
		mux.Post("/readFhx", app.ReadFhx)
	})

	// FHX Operationen Routes
	mux.Route("/operation", func(mux chi.Router) {
		mux.Get("/{plantId}", app.OperationPage)
		mux.Get("/details/{opplantId}", app.OperationDetails)
	})

	mux.NotFound(app.NotFound)

	return mux
}
