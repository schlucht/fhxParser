package models

import (
	"database/sql"
	"log"
	"os"
)

// for Database connection
type DBModel struct {
	DB       *sql.DB
	infoLog  *log.Logger
	errorLog *log.Logger
}

// New Models return a Model
func NewModel(db *sql.DB) DBModel {
	infoLog := log.New(os.Stdout, "\x1b[32mINFO_DB:\x1b[0m\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "\x1b[31mError_DB:\x1b[0m\t", log.Ldate|log.Ltime|log.Lshortfile)
	return DBModel{
		DB:       db,
		infoLog:  infoLog,
		errorLog: errorLog,
	}
}
