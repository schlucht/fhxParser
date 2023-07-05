package main

import (
	"encoding/json"
	"net/http"

	"github.com/schlucht/fhxreader/internal/parser"
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

// Liest die Datei
func (app *application) HomeReadFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Fatal(err)
	}
	fhxTxt := r.FormValue("text-output")
	// log.Println(len(strings.Split(fhxTxt, "\n")))
	// Parsen der eingelesen strings
	fhx, err := parser.NewFhxString(fhxTxt)
	if err != nil {
		app.errorLog.Fatal(err)
	}
	// log.Println(fhx)
	f, err := json.MarshalIndent(fhx, "", "   ")
	if err != nil {
		app.errorLog.Fatal(err)
	}
	j := jsonResponse{
		OK:      true,
		Message: "geklappt",
		Content: string(f),
		ID:      1,
	}

	b, err := json.Marshal(j)
	if err != nil {
		app.errorLog.Fatal(err)
	}
	// w.Header().Set("Content-Type", "application/json")
	var data = make(map[string]interface{})
	data["data"] = string(b)
	// w.Write([]byte(b))
	// fmt.Printf("%v", string(b))
	if app.renderTemplate(w, r, "fhxData", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}
