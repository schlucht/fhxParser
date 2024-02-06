package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/schlucht/fhxreader/internal/parser"
)

type antwort struct {
	Id int `json:"id"`
}

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

var j = jsonResponse{
	OK:      false,
	Message: "",
	Content: "{}",
	ID:      999,
}

// gibt alle Abteilungen zurück
func (app *application) AllPlants(w http.ResponseWriter, r *http.Request) {
	allPlants, err := app.DB.LoadAllPlants()
	if err != nil {
		app.errorLog.Printf("Fehler beim auslesen der Anlage %v", err)
		app.badRequest(w, r, err)
	}

	s, err := json.Marshal(allPlants)
	if err != nil {
		app.errorLog.Printf("Fehler beim parsen der Anlage!")
		app.badRequest(w, r, err)
	}

	if len(s) == 0 {
		j.OK = false
		j.Message = "Keine Anlagen gefunden"
		j.Content = "{}"
	} else {
		j.OK = true
		j.Message = "Anlagen gefunden"
		j.Content = string(s)
	}

	app.writeJSON(w, http.StatusOK, j)
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

// Procedure in der Datenbank speichern
func (app *application) saveProcedure(fhx parser.Fhx, plantId int) (string, error) {

	return "", nil
}
