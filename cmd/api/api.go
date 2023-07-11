package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/schlucht/fhxreader/internal/driver"
	"github.com/schlucht/fhxreader/internal/parser"
)

const version = "1.0.0"
const cssVersion = "1"

type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	app.infoLog.Printf("Back Server run on Port: %v on mode: %s\n", app.config.port, app.config.env)
	return srv.ListenAndServe()
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 5001, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application enviroment { develompen | production}")
	flag.StringVar(&cfg.db.dsn, "dsn", "schmidschluch4:Schlucht6@tcp(db8.hostpark.net)/schmidschluch4", "DB connect String")
	flag.StringVar(&cfg.api, "api", "http://localhost:5001", "URL to API")
	//flag.StringVar(&cfg.api, "api", "https://5001-schlucht-fhxparser-zz2ewe38uk4.ws-eu101.gitpod.io", "URL to API")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := driver.OpenDB(cfg.db.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}

func (app *application) loadOperations(txt string) ([]parser.Fhx, error) {

	ops, err := parser.NewFhxString(txt)

	fmt.Println("**************************************")
	fmt.Println("Name:", ops[0].OPs[0].UnitName)
	fmt.Println("**************************************")
	for _, o := range ops[0].OPs[0].Parameters {
		fmt.Println(o.Name, o.Value)
	}

	if err != nil {
		return nil, err
	}

	return ops, nil
}
