package main

import (
	"database/sql"
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

type config struct {
	port int
	env  string
	db   struct {
		dsn    string
		driver string
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

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type envelope map[string]interface{}

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
	flag.StringVar(&cfg.db.dsn, "dsn", "ots:fhxdb@tcp(localhost:3306)/fhxdb?parseTime=true", "DB connect String")
	// flag.StringVar(&cfg.db.dsn, "dsn", "schmidschluch4:Schlucht6@tcp(db8.hostpark.net)/schmidschluch4?parseTime=true", "DB connect String")

	flag.Parse()

	infoLog := log.New(os.Stdout, "\n\x1b[32mINFO:\x1b[0m\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "\n\x1b[31mERROR:\x1b[0m\t", log.Ldate|log.Ltime|log.Lshortfile)

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	var conn *sql.DB
	var err error

	conn, err = driver.MySqlOpenDB(cfg.db.dsn)
	if err != nil {
		errorLog.Printf("%v", err)
	}

	mod := models.NewModel(conn)
	defer conn.Close()

	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       version,
		Session:       session,
		DB:            mod,
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
