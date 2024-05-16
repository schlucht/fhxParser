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

// Wrapper for all Models
type Models struct {
	DB DBModel
}

// New Models return a Model
func NewModel(db *sql.DB) Models {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return Models{
		DB: DBModel{
			DB:       db,
			infoLog:  infoLog,
			errorLog: errorLog,
		},
	}
}
