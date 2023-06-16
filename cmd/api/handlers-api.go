package main

import (
	"encoding/json"
	"net/http"
)

type fhxFileLoad struct {
	FileName string `json:"file-name"`
	FileText string `json:"file-text"`
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Content string `json:"content"`
	ID      int    `json:"id"`
}

func (app *application) HomeReadFile(w http.ResponseWriter, r *http.Request) {

	var payload fhxFileLoad
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	j := jsonResponse{
		OK:      true,
		Message: "",
		Content: "",
		ID:      10,
	}
	out, err := json.MarshalIndent(j, "", "   ")
	if err != nil {
		app.errorLog.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}
