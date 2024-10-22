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
