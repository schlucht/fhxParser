package main

import (
	"log"
	"os"

	"github.com/schlucht/fhxreader/fhx/fhxFactory"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {

	infoLog := log.New(os.Stdout, "\033[0;34mINFO: \033[0m \t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "\033[0;31mERROR: \033[0m \t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	err := app.loadFile()
	if err != nil {
		app.errorLog.Fatal("no data load")
	}
}

func (app *application) loadFile() error {
	ff := fhxFactory.Load("UP_Q2000_START")
	app.infoLog.Println(ff.File)
	return nil
}
