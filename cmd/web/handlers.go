package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Content string `json:"content"`
	ID      int    `json:"id"`
}

type loadFile struct {
	FileName string `json:"file-name"`
	FileText string `json:"file-text"`
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "home", nil); err != nil {
		app.errorLog.Println(err)
	}
}
func (app *application) HomeReadFile(w http.ResponseWriter, r *http.Request) {

	j := jsonResponse{
		OK:      true,
		Message: "geklappt",
		Content: "kommt",
		ID:      1,
	}
	json.MarshalIndent(j, "", "   ")
	if err := app.renderTemplate(w, r, "home", nil); err != nil {
		app.errorLog.Println(err)
	}
}
