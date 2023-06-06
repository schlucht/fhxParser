package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/schlucht/fhxreader/fhx-app/config"
	"github.com/schlucht/fhxreader/fhx-app/driver"
	"github.com/schlucht/fhxreader/fhx-app/handlers"
	"github.com/schlucht/fhxreader/fhx-app/render"
)

type Server struct {
	Port string
}

var app config.AppConfig
var session *scs.SessionManager

func run() (*driver.DB, error) {
	name := "schmidschluch4"
	host := "db8.hostpark.net"
	pw := "Schlucht6"
	connectString := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		name, pw, host, name,
	)
	log.Println("Start Connected to database...")

	db, err := driver.ConnectSQL(connectString)
	if err != nil {
		return nil, err
	}

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	app.InProduction = false

	tc, err := render.CreateTemplateCache()
	if err != nil {
		return nil, errors.New("can not create template")
	}
	app.TemplateCatche = tc
	app.UseCache = false

	return db, nil
}

// Startet den Web- und den Datenbank Server
func (m *Server) Start() {
	m.Port = ":8080"

	db, err := run()
	if err != nil {
		log.Fatalf("Can not open DB %v", err)
	}
	log.Println("Ende Connected to database...")
	defer db.SQL.Close()

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandler(repo)
	render.NewRenderer((&app))

	srv := &http.Server{
		Addr:    m.Port,
		Handler: routes(&app),
	}

	fmt.Println("Server run of localhost:" + m.Port)

	srv.ListenAndServe()
}
