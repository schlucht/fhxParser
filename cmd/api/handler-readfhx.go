package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/schlucht/fhxreader/internal/helpers"
	"github.com/schlucht/fhxreader/internal/parser"
)

// Liest einen FHX Text ein. Es muss der Text und eine ID für eine Anlage vorhanden sein.

func (app *application) ReadFhx(w http.ResponseWriter, r *http.Request) {
	var fhxJson struct {
		FileText string `json:"text"`
		FileName string `json:"name"`
		PlantId  string `json:"plant_id"`
	}
	var payload jsonResponse
	err := app.readJSON(w, r, &fhxJson)
	if err != nil {
		app.badRequest(w, err, "ReadFhx: readJson", http.StatusInternalServerError)
		return
	}

	fhx, err := parser.NewFhxString(fhxJson.FileText)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Fehler im FHX Parser")
		app.badRequest(w, err, "ReadFhx: NewFhxString", http.StatusInternalServerError)
		return
	}

	for _, f := range fhx {
		if f.UnitType == "OPERATION" {
			helpers.SaveJSON("assets/files/operation.json", helpers.PrintJson(f))
			// Speichern der Operation
			err := app.SaveOperation(f, uuid.MustParse(fhxJson.PlantId))
			if err != nil {
				app.errorLog.Printf("%v, %s", err, "Fehler im FHX Parser")
				app.badRequest(w, err, "ReadFhx: NewFhxString", http.StatusInternalServerError)
				return
			}
			payload.Message = "OK, Save Operations"
			payload.Error = true
		} else if f.UnitType == "UNIT_PROCEDURE" {
			helpers.SaveJSON("assets/files/units.json", helpers.PrintJson(f))
			// Speichern der Unit
			err = app.SaveAllUnits(f, uuid.MustParse(fhxJson.PlantId))
			if err != nil {
				app.badRequest(w, err, "ReadFhx: NewFhxString", http.StatusInternalServerError)
				return
			}
			payload.Message = "OK, Save Units"
			payload.Error = false
		} else if f.UnitType == "PROCEDURE" {
			helpers.SaveJSON("assets/files/procedure.json", helpers.PrintJson(f))
			payload.Message = "OK, Save Recipes"
			payload.Error = false
		}
	}

	w.Header().Set("Content-Type", "application/json")
	app.writeJSON(w, http.StatusOK, payload)
}
