package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/schlucht/fhxreader/internal/helpers"
	"github.com/schlucht/fhxreader/internal/parser"
)

// Liest einen FHX Text ein. Es muss der Text und eine ID für eine Anlage vorhanden sein.
func (app *application) ReadFhx(w http.ResponseWriter, r *http.Request) {

	f, err := io.ReadAll(r.Body)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Keine Daten vorhanden")
		app.badRequest(w, r, err)
		return
	}

	var fhxJson = fhxFileLoad{}
	err = json.Unmarshal(f, &fhxJson)
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, r, err)
		return
	}

	fhx, err := parser.NewFhxString(fhxJson.FileText)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Fehler im FHX Parser")
		j.OK = false
		j.Message = fmt.Sprintf("%v", err)
		app.writeJSON(w, http.StatusOK, j)
		return
	}

	for _, f := range fhx {
		if f.UnitType == "OPERATION" {
			helpers.SaveJSON("assets/files/operation.json", helpers.PrintJson(f))
			j.Message = "OK, Save Operations"
			j.OK = true
		} else if f.UnitType == "UNIT_PROCEDURE" {
			helpers.SaveJSON("assets/files/units.json", helpers.PrintJson(f))
			j.Message = "OK, Save Units"
			j.OK = true
		} else if f.UnitType == "PROCEDURE" {
			helpers.SaveJSON("assets/files/procedure.json", helpers.PrintJson(f))
			j.Message = "OK, Save Recipes"
			j.OK = true
		}
	}

	w.Header().Set("Content-Type", "application/json")
	app.writeJSON(w, http.StatusOK, j)
}
