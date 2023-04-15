package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	port int
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	var cfg config
	cfg.port = 8081
	infoLog := log.New(os.Stdout, "\033[0;34mINFO: \033[0m \t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "\033[0;31mERROR: \033[0m \t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	err := app.serve()
	if err != nil {
		errorLog.Fatal(err)
	}

	// p := "./files/q2000.fhx"
	// fhxFactory.SaveFhxFile(p)
}

func (app *application) serve() error {
	app.infoLog.Println("API listion on port", app.config.port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	return srv.ListenAndServe()
}
