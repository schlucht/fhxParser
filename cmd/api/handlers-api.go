package main

import (
	"io"
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

func (app *application) ReadFhx(w http.ResponseWriter, r *http.Request) {

	f, err := io.ReadAll(r.Body)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	txt := string(f)
	_, err = app.insertOperations(txt)
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, r, err)
		return
	}
	// app.infoLog.Println(fhx)
	j := jsonResponse{
		OK:      true,
		Message: "Hochladen hat geklappt",
		Content: "Alles gut",
		ID:      10,
	}
	w.Header().Set("Content-Type", "application/text")
	app.writeJSON(w, http.StatusOK, j)
}
