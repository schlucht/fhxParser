package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
		Message: "Hochladen hat geklappt",
		Content: "",
		ID:      999,
	}

	f, err := io.ReadAll(r.Body)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	var fhxJson fhxFileLoad = fhxFileLoad{}
	err = json.Unmarshal(f, &fhxJson)
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, r, err)
		return
	}
	app.infoLog.Println(fhxJson.FileName, fhxJson.PlantId)

	doubleUp, err := app.insertOperations(fhxJson.FileText, int(fhxJson.PlantId))
	if err != nil {
		app.errorLog.Println(err)
		return
	}
	if len(doubleUp) > 0 {
		s := strings.Join(doubleUp, ",")
		j.OK = false
		j.Message = fmt.Sprintf("Dopplete UP's: %s", s)
	}
	app.infoLog.Println(j)

	w.Header().Set("Content-Type", "application/text")
	app.writeJSON(w, http.StatusOK, j)
}

// gibt alle Abteilungen zurück
func (app *application) AllPlants(w http.ResponseWriter, r *http.Request) {
	allPlants, err := app.DB.LoadAllPlants()
	app.infoLog.Println(allPlants)
	if err != nil {
		app.badRequest(w, r, err)
	}
	app.writeJSON(w, http.StatusOK, allPlants)
}
