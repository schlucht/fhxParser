package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/schlucht/fhxreader/fhx-app/config"
	"github.com/schlucht/fhxreader/fhx-app/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
