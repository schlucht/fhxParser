package main

import (
	"net/http"
)

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Content string `json:"content"`
	ID      int    `json:"id"`
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "home", &templateData{
		FileName: "/",
	}); err != nil {
		app.errorLog.Println(err)
	}
}

// Ladet alle Anlage zur Auswahl in die Seite
