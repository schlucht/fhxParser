package main

import (
	"fmt"
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
	ff := fhxFactory.Load("./assets/files/q2000.fhx")
	fmt.Println(ff.File)
	for _, f := range ff.Fhx {
		for _, u := range f.Units {
			fmt.Println(u.UnitPosition)
			for _, p := range u.Procedures {
				fmt.Printf("\t%s\n", p.Name)
				for _, pa := range p.Parameters {
					fmt.Printf("\t\t%s\t|Value: %s\n", pa.Name, pa.Value.Cv)
				}
			}
		}
	}
	return nil
}
