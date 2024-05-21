package main

import (
	"net/http"
)

func (app *application) PlantPage(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "plant", &templateData{}); err != nil {
		app.errorLog.Println(err)
	}
}
