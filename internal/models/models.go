package models

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"
	"time"
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

func (m *DBModel) GetQueryText(path string) (string, error) {
	var sql = ""
	data, err := os.ReadFile(path)
	if err != nil {
		return "", nil
	}
	if len(data) == 0 {
		return "", errors.New("no data in sql file")
	}
	sql = string(data)
	return sql, nil
}

// Funtion zum Erstellen einer Tabelle.
// Parameter: SQL String zum erstellen der Tabelle
// Return: error
func (m *DBModel) CreateTable(sql string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := m.DB.ExecContext(ctx, sql)
	if err != nil {
		return err
	}
	return nil
}
