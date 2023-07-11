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

	var payload fhxFileLoad = fhxFileLoad{}

	f, err := io.ReadAll(r.Body)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	payload.FileText = string(f)

	j := jsonResponse{
		OK:      true,
		Message: "Hochladen hat geklappt",
		Content: "Alles gut",
		ID:      10,
	}
	app.writeJSON(w, http.StatusOK, j)
	// out, err := json.MarshalIndent(j, "", "   ")
	// if err != nil {
	// 	app.errorLog.Println(err)
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(out)
}
