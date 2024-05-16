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
const frontend_url = "127.0.0.1:5100"

var session *scs.SessionManager

type config struct {
	port int
	env  string
	db   struct {
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
	app.infoLog.Printf("Back Server run on Port: %v on mode: %s\n", app.config.port, app.config.env)
	return srv.ListenAndServe()
}

func main() {
	var cfg config
	tc := make(map[string]*template.Template)
	flag.IntVar(&cfg.port, "port", 5101, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application enviroment { develompen | production}")
	flag.StringVar(&cfg.db.dsn, "dsn", "schmidschluch4:Schlucht6@tcp(db8.hostpark.net)/schmidschluch4?parseTime=true", "DB connect String")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	session = scs.New()
	session.Lifetime = 24 * time.Hour

	conn, err := driver.OpenDB(cfg.db.dsn)

	if err != nil {
		errorLog.Printf("%v", err)
	}
	defer conn.Close()

	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       version,
		Session:       session,
		DB:            models.DBModel{DB: conn},
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
