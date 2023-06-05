package handlers

import (
	"net/http"

	"github.com/schlucht/fhxreader/fhx-app/models"
	"github.com/schlucht/fhxreader/fhx-app/render"
)

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	stringMap["title"] = "About"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	}, r)
}
