package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/schlucht/fhxreader/internal/driver"
	"github.com/schlucht/fhxreader/internal/models"
)

const version = "1.0.0"
var session *scs.SessionManager

const serverURL = "127.0.0.1:5101"
const frontendURL = "127.0.0.1:5100"

type config struct {
	port     int
	env      string
	api      string
	frontend string
	db       struct {
		dsn string
	}
}

type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
	Session       *scs.SessionManager
	DB            models.DBModel
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
	app.infoLog.Printf("Server run on Port: %v in mode: %s\n", app.config.port, app.config.env)
	return srv.ListenAndServe()
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 5100, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application enviroment { develompen | production}")
	flag.StringVar(&cfg.db.dsn, "dsn", "schmidschluch4:Schlucht6@tcp(db8.hostpark.net)/schmidschluch4", "DB connect String")
	flag.StringVar(&cfg.api, "api", serverURL, "URL to API")
	flag.StringVar(&cfg.frontend, "frontend", frontendURL, "URL to Frontend")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := driver.OpenDB(cfg.db.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	session = scs.New()
	session.Lifetime = time.Hour * 24

	tc := make(map[string]*template.Template)

	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       version,
		Session:       session,
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
