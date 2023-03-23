package main

import (
	"log"
	"os"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{

		infoLog:  infoLog,
		errorLog: errorLog,
	}

	err := app.loadData()
	if err != nil {
		app.errorLog.Fatal("no data load")
	}
}

func (app *application) loadData() error {
	app.infoLog.Println("fhx Works")
	return nil
}
