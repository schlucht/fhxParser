package main

import "net/http"

type antwort struct {
	Id int `json:"id"`
}

type jsonResponse struct {
	OK      bool        `json:"ok"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
	ID      int         `json:"id"`
}

var j = jsonResponse{
	OK:      false,
	Message: "",
	Content: "",
	ID:      999,
}

// gibt alle Abteilungen zurück

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "home", &templateData{}, "loadFile"); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) NotFound(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "notFound", &templateData{}); err != nil {
		app.errorLog.Println(err)
	}
}
