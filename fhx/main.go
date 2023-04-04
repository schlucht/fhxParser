package main

import (
	"encoding/json"
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
	ff, err := fhxFactory.Load("./assets/files/q2000.fhx")
	if err != nil {
		app.errorLog.Println(err)
	}
	err = ff.SaveUnits()
	if err != nil {
		app.errorLog.Println(err)
	}
	return nil
}

func PrintJson(ff *fhxFactory.FhxFactory) (string, error) {
	data, err := json.Marshal(ff.Fhx)
	if err != nil {
		return "", err
	}

	return string(data), nil

}
