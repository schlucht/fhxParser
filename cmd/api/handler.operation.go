package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (app *application) allOperations(w http.ResponseWriter, r *http.Request) {

	plantId := chi.URLParam(r, "plantId")
	if plantId == "" {
		app.errorLog.Println(plantId)
		app.badRequest(w, nil, "allOperations: plantId", http.StatusBadRequest)
		return
	}	

	pid := uuid.MustParse(plantId)
	all, err := app.DB.OpFromPlantID(pid)
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, err, "allOperations: GetAllOperations", http.StatusInternalServerError)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "OK",
		Data:    all,
	}

	app.writeJSON(w, http.StatusOK, payload)

}
