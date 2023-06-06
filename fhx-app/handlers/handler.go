package handlers

import (
	"github.com/schlucht/fhxreader/fhx-app/config"
	"github.com/schlucht/fhxreader/fhx-app/driver"
	"github.com/schlucht/fhxreader/repository"
	"github.com/schlucht/fhxreader/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

func NewRepo(m *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: m,
		DB:  dbrepo.NewMySqlRepo(db.SQL, m),
	}
}

func NewHandler(r *Repository) {
	Repo = r
}
