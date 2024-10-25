package main

import (
	"net/http"
)

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

// Suggested code may be subject to a license. Learn more: ~LicenseLog:244431160.
func (app *application) SavePlant(w http.ResponseWriter, r *http.Request) {

	err := app.DB.CreatePlantTable()
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, err, "SavePlant: CreatePlantTable", http.StatusInternalServerError)
		return
	}

	var requestPayload struct {
		Plant string `json:"plant"`		
	}

	if err := app.readJSON(w, r, &requestPayload); err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, err, "SavePlant: readJSON", http.StatusBadRequest)
		return
	}

	err = app.DB.CreateNewPlant(requestPayload.Plant)
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, err, "SavePlant: CreateNewPlant", http.StatusInternalServerError)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
		Data:    envelope{"plant": requestPayload.Plant},
	}

	if err = app.writeJSON(w, http.StatusOK, payload); err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, err, "SavePlant: writeJSON", http.StatusInternalServerError)
		return
	}

	
}