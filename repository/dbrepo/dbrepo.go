package dbrepo

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/schlucht/fhxreader/fhx-app/config"
)

type mysqlDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewMySqlRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &mysqlDBRepo{
		App: a,
		DB:  conn,
	}
}
