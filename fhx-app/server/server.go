package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/schlucht/fhxreader/fhx-app/config"
	"github.com/schlucht/fhxreader/fhx-app/handlers"
	"github.com/schlucht/fhxreader/fhx-app/render"
)

type Server struct {
	Port string
}

var app config.AppConfig
var session *scs.SessionManager

func (m *Server) Start() {

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	app.Session = session

	app.InProduction = false

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can not create Template")
	}
	app.TemplateCatche = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplate((&app))
	mux := routes(&app)

	m.Port = ":8080"
	fmt.Println("Server run of localhost:" + m.Port)
	http.ListenAndServe(m.Port, mux)
}
