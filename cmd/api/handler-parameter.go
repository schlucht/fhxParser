package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// Alle Parameter der jeweiligen Operation einlesen
func (app *application) getParamsFromOPId(w http.ResponseWriter, r *http.Request) {

	f, err := io.ReadAll(r.Body)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Keine Daten vorhanden")
		app.badRequest(w, r, err)
		return
	}

	var re antwort
	err = json.Unmarshal(f, &re)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Id der OP fehlt konnte nicht geparst werden")
		app.badRequest(w, r, err)
		return
	}
	if int(re.Id) > 0 {
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
		if len(s) == 0 {
			j.OK = false
			j.Message = "Keine Parameter gefunden"
			j.Content = "{}"
		} else {
			j.OK = true
			j.Content = string(s)
			j.Message = "Daten OK"
		}
	} else {
		j.OK = false
		j.Message = "Keine ID vorhanden!"
		j.Content = "{}"
	}

	w.Header().Set("Content-Type", "application/json")
	app.writeJSON(w, http.StatusOK, j)
}
