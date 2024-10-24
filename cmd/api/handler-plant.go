package main

import (
	"net/http"
)

// Alle Anlagen aus der Datenbank an Frontend übergeben
func (app *application) AllPlants(w http.ResponseWriter, r *http.Request) {

	all, err := app.DB.GetPlants()
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, err, "AllPlants: GetPlants", http.StatusInternalServerError)
		return
	}
	payload := jsonResponse{
		Error:   false,
		Message: "success",
		Data:    envelope{"plants": all},
	}

	if err = app.writeJSON(w, http.StatusOK, payload); err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, err, "AllPlants: writeJSON", http.StatusInternalServerError)
		return
	}
}

// Anlage in der Datenbank speichern
func (app *application) SavePlant(w http.ResponseWriter, r *http.Request) {
	err := app.DB.CreatePlantTable()
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, err, "SavePlant: CreatePlantTable", http.StatusNotModified)
		return
	}
	type credentials struct {
		Plant string `json:"plant"`
	}
	var creds credentials
	// var payload jsonResponse

	err = app.readJSON(w, r, &creds)
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, err, "SavePlant: Read Plant", http.StatusInternalServerError)
		return
	}
	err = app.DB.NewPlant(creds.Plant)
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, err, "SavePlant: InsertPlant", http.StatusInternalServerError)
		return
	}

	app.infoLog.Println(creds.Plant)

}
