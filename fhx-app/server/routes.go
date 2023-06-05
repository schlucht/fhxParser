package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/schlucht/fhxreader/fhx-app/config"
	"github.com/schlucht/fhxreader/fhx-app/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Post("/readFile", handlers.Repo.ReadFile)

	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./fhx-app/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
