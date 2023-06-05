package handlers

import "github.com/schlucht/fhxreader/fhx-app/config"

var Repo *Repository
type Repository struct {
	App *config.AppConfig
}

func NewRepo(m *config.AppConfig) *Repository {
	return &Repository{
		App: m,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}
