package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// Alle Parameter der jeweiligen Operation einlesen
func (app *application) getParamsFromOPId(w http.ResponseWriter, r *http.Request) {
	j := jsonResponse{
		OK:      true,
		Message: "Daten konnten nicht gespeichert werden",
		Content: "",
		ID:      999,
	}

	f, err := io.ReadAll(r.Body)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Keine Daten vorhanden")
		app.badRequest(w, r, err)
		return
	}

	var re antwort
	err = json.Unmarshal(f, &re)
	app.infoLog.Println(re.Id)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Id der OP fehlt konnte nicht geparst werden")
		app.badRequest(w, r, err)
		return
	}

	opts, err := app.GetParamFromOPId(int(re.Id))
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Fehler beim Laden von Parameter")
		app.badRequest(w, r, err)
		return
	}

	s, err := json.Marshal(opts)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Parameter konnte nicht geparst werden")
		app.badRequest(w, r, err)
		return
	}
	j.Content = string(s)

	w.Header().Set("Content-Type", "application/json")
	app.writeJSON(w, http.StatusOK, j)
}
