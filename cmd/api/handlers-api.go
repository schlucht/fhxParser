package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/schlucht/fhxreader/internal/parser"
)

type fhxFileLoad struct {
	FileText string `json:"text"`
	FileName string `json:"name"`
	PlantId  int    `json:"plant_id"`
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Content string `json:"content"`
	ID      int    `json:"id"`
}

// Liest einen FHX Text ein. Es muss der Text und eine ID für eine Anlage vorhanden sein.
func (app *application) ReadFhx(w http.ResponseWriter, r *http.Request) {
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

	var fhxJson = fhxFileLoad{}
	err = json.Unmarshal(f, &fhxJson)
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, r, err)
		return
	}

	fhx, err := parser.NewFhxString(fhxJson.FileText)
	if err != nil {
		app.errorLog.Println(err)
		j.OK = false
		j.Message = fmt.Sprintf("%v", err)
		app.writeJSON(w, http.StatusOK, j)
		return
	}
	app.infoLog.Println("OP: ", fhx[0].UnitType)
	for _, f := range fhx {
		if f.UnitType == "OPERATION" {
			msg, err := app.saveOperations(f, int(fhxJson.PlantId))
			if err != nil {
				j.OK = false
				j.Message = fmt.Sprintf("%v", err)
				app.writeJSON(w, http.StatusOK, j)
				return
			}
			j.Message = msg
			j.OK = true
		}
	}

	w.Header().Set("Content-Type", "application/json")
	app.writeJSON(w, http.StatusOK, j)
}

// gibt alle Abteilungen zurück
func (app *application) AllPlants(w http.ResponseWriter, r *http.Request) {
	allPlants, err := app.DB.LoadAllPlants()
	// app.infoLog.Println(allPlants)
	if err != nil {
		app.badRequest(w, r, err)
	}
	app.writeJSON(w, http.StatusOK, allPlants)
}

// Operation in der Datenbank speichern
func (app *application) saveOperations(fhx parser.Fhx, plantId int) (string, error) {
	doubleUp, err := app.insertOperations(fhx, plantId)
	if err != nil {
		return "", err
	}
	if len(doubleUp) > 0 {
		s := strings.Join(doubleUp, ",")
		return fmt.Sprintf("Dopplete UP's: %s", s), nil
	}
	return "", nil
}
