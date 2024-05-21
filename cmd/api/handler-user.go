package main

import (
	"net/http"
)

func (app *application) UserPage(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "users", &templateData{}); err != nil {
		app.errorLog.Println(err)
	}
}
