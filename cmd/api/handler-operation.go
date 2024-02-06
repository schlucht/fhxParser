package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// Alle Operationen der jeweiligen Operation einlesen
func (app *application) GetOperations(w http.ResponseWriter, r *http.Request) {

	f, err := io.ReadAll(r.Body)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Keine Daten vorhanden")
		app.badRequest(w, r, err)
		return
	}
	var re antwort

	err = json.Unmarshal(f, &re)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Betrieb fehlt konnte nicht geparst werden")
		app.badRequest(w, r, err)
		return
	}

	opts, err := app.AllOperations(int(re.Id))
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Fehler beim Laden von Operation")
		app.badRequest(w, r, err)
		return
	}

	s, err := json.Marshal(opts)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Operation konnte nicht geparst werden")
		app.badRequest(w, r, err)
		return
	}

	if opts.Count == 0 {
		j.OK = false
		j.Content = "{}"
		j.Message = "Keine Operation gefunden"
	} else {
		j.OK = true
		j.Message = "Daten OK"
		j.Content = string(s)
	}

	w.Header().Set("Content-Type", "application/json")
	app.writeJSON(w, http.StatusOK, j)
}

// Alle Operationen der jeweiligen Operation einlesen
func (app *application) GetOperationFromId(w http.ResponseWriter, r *http.Request) {
	f, err := io.ReadAll(r.Body)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Keine Daten vorhanden")
		app.badRequest(w, r, err)
		return
	}
	var re antwort
	err = json.Unmarshal(f, &re)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Betrieb fehlt konnte nicht geparst werden")
		app.badRequest(w, r, err)
		return
	}

	opt, err := app.GetOpFromId(int(re.Id))
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Fehler beim Laden von Operation")
		app.badRequest(w, r, err)
		return
	}

	s, err := json.Marshal(opt)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Operation konnte nicht geparst werden")
		app.badRequest(w, r, err)
		return
	}

	if len(s) == 0 {
		j.OK = false
		j.Message = "Keine Daten vorhanden!"
		j.Content = "{}"
	} else {
		j.OK = true
		j.Message = "Daten OK"
		j.Content = string(s)
	}

	w.Header().Set("Content-Type", "application/json")
	app.writeJSON(w, http.StatusOK, j)
}
