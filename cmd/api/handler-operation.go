package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (app *application) OperationPage(w http.ResponseWriter, r *http.Request) {
	plantId := chi.URLParam(r, "plantId")

	uuid, err := uuid.Parse(plantId)
	if err != nil {
		app.badRequest(w, err, "UUID nicht IO", http.StatusInternalServerError)
		return
	}

	plants, err := app.DB.OpFromPlantID(uuid)
	if err != nil {
		app.badRequest(w, err, "GetPlantsOp", http.StatusNoContent)
		return
	}

	// Daten an das Frontend übergeben
	data := make(map[string]interface{})
	data["ops"] = plants

	if err := app.renderTemplate(w, r, "operation", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) OperationDetails(w http.ResponseWriter, r *http.Request) {
	// TODO
	opplantId := chi.URLParam(r, "opplantId")
	app.infoLog.Println(opplantId)
	// var body struct {
	// 	Id string `json:"id"`
	// }
	// err := app.readJSON(w, r, &body)
	// if err != nil {
	// 	app.badRequest(w, r, err, "OperationDetail: readJson")
	// 	return
	// }

	data := make(map[string]interface{})
	params, err := app.opDetailsFromOp(opplantId)
	if err != nil {
		app.badRequest(w, err, "OperationDetail: opDetailsFromOp", http.StatusNoContent)
		return
	}

	data["params"] = params
	// if err = app.writeJSON(w, http.StatusOK, params); err != nil {
	// 	app.errorLog.Println(err)
	// }
	if err := app.renderTemplate(w, r, "operation", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
	
}
