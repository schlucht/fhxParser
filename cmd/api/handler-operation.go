package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type antwort struct {
	PlantId int `json:"plantId"`
}

// Alle Operationen der jeweiligen Operation einlesen
func (app *application) GetOperations(w http.ResponseWriter, r *http.Request) {
	j := jsonResponse{
		OK:      true,
		Message: "Daten konnten nicht gespeichert werden",
		Content: "",
		ID:      999,
	}
	f, err := io.ReadAll(r.Body)
	if err != nil {
		app.infoLog.Printf("%s", "Keine Daten vorhanden")
		app.badRequest(w, r, err)
		return
	}
		
	var re antwort
	err = json.Unmarshal(f, &re)
	if err != nil {
		app.infoLog.Printf("%s", "Betrieb fehlt konnte nicht geparst werden")
		app.badRequest(w, r, err)
		return
	}
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, r, err)
		return
	}
	opts, err := app.AllOperations(int(re.PlantId))

	if err != nil {
		app.infoLog.Printf("%s", "Fehler beim Laden von Operation")
		app.badRequest(w, r, err)
		return
	}

	s, err := json.Marshal(opts)
	if err != nil {
		app.infoLog.Printf("%s", "Operation konnte nicht geparst werden")
		app.badRequest(w, r, err)
		return
	}
	j.Content = string(s)

	w.Header().Set("Content-Type", "application/json")
	app.writeJSON(w, http.StatusOK, j)
}
