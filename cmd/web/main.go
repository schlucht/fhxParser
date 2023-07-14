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
const cssVersion = "1"

var session *scs.SessionManager

const gitpodServer = "https://5101-schlucht-fhxparser-zz2ewe38uk4.ws-eu101.gitpod.io"
const homeServer = "http://localhost:5101"

type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
	stripe struct {
		secret string
		key    string
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
	flag.StringVar(&cfg.api, "api", gitpodServer, "URL to API")

	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

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
