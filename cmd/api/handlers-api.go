package main

import (
	"fmt"
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
