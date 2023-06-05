package handlers

import (
	"net/http"

	"github.com/schlucht/fhxreader/fhx-app/models"
	"github.com/schlucht/fhxreader/fhx-app/render"
)

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{}, r)
}
