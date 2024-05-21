package main

import (
	"net/http"
)

func (app *application) FhxPage(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "fhx", &templateData{}); err != nil {
		app.errorLog.Println(err)
	}
}
