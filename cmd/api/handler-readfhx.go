package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/schlucht/fhxreader/internal/helpers"
	"github.com/schlucht/fhxreader/internal/parser"
)

// Liest einen FHX Text ein. Es muss der Text und eine ID für eine Anlage vorhanden sein.

func (app *application) FhxPage(w http.ResponseWriter, r *http.Request) {
	plants, err := app.LoadPlants()
	if err != nil {
		app.errorLog.Println(err)
	}
	if len(plants) > 0 {
		// Daten an das Frontend übergeben
		data := make(map[string]interface{})
		data["plants"] = plants

		if err := app.renderTemplate(w, r, "fhx", &templateData{
			Data: data,
		}); err != nil {
			app.errorLog.Println(err)
		}
	} else {
		if err := app.renderTemplate(w, r, "plant", &templateData{}); err != nil {
			app.errorLog.Println(err)
		}
	}
	// if err := app.renderTemplate(w, r, "fhx", &templateData{}); err != nil {
	// 	app.errorLog.Println(err)
	// }
}

func (app *application) ReadFhx(w http.ResponseWriter, r *http.Request) {
	var fhxJson struct {
		FileText string `json:"text"`
		FileName string `json:"name"`
		PlantId  string `json:"plant_id"`
	}
	err := app.readJSON(w, r, &fhxJson)
	if err != nil {
		app.badRequest(w, r, err, "ReadFhx: readJson")
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
			// Speichern der Operation
			err := app.SaveOperation(f, uuid.MustParse(fhxJson.PlantId))
			if err != nil {
				app.errorLog.Printf("%v, %s", err, "Fehler im FHX Parser")
				j.OK = false
				j.Message = fmt.Sprintf("%v", err)
				app.writeJSON(w, http.StatusOK, j)
				return
			}
			j.Message = "OK, Save Operations"
			j.OK = true
		} else if f.UnitType == "UNIT_PROCEDURE" {
			helpers.SaveJSON("assets/files/units.json", helpers.PrintJson(f))
			// Speichern der Unit
			err = app.SaveUnit(f, uuid.MustParse(fhxJson.PlantId))
			if err != nil {
				j.OK = false
				j.Message = fmt.Sprintf("%v", err)
				app.writeJSON(w, http.StatusOK, j)
				return
			}
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
