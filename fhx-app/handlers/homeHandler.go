package handlers

import (
	"log"
	"net/http"

	"github.com/schlucht/fhxreader/fhx-app/models"
	"github.com/schlucht/fhxreader/fhx-app/render"
	"github.com/schlucht/fhxreader/fhx-parser/parser"
)

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	}, r)
}

func (m *Repository) ReadFile(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Printf("Fehler 1: %v", err)
		return
	}
	txt := r.Form.Get("textout")
	fhx, err := parser.NewFhxString(txt)
	if err != nil {
		log.Printf("Fehler: %v", err)
	}
	log.Println(len(fhx[0].Units))
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{}, r)
}
