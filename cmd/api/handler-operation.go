package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (app *application) OperationPage(w http.ResponseWriter, r *http.Request) {
	plantId := chi.URLParam(r, "plantId")
	app.infoLog.Println(plantId)
	uuid, err := uuid.Parse(plantId)
	if err != nil {
		app.badRequest(w, r, err, "UUID nicht IO")
		return
	}
	plants, err := app.DB.OpFromPlantID(uuid)
	if err != nil {
		app.badRequest(w, r, err, "GetPlantsOp")
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

func (app *application) OperationSave(w http.ResponseWriter, r *http.Request) {
	// TODO
	plantId := chi.URLParam(r, "plantId")
	opPlantId := chi.URLParam(r, "opPlantId")
	app.infoLog.Println(plantId, opPlantId)

	var body struct {
		Id string `json:"id"`
	}
	// err := app.readJSON(w, r, &body)
	// if err != nil {
	// 	app.badRequest(w, r, err, "OperationSave: readJson")
	// 	return
	// }
	data := make(map[string]interface{})
	data["body"] = body
	if err := app.renderTemplate(w, r, "operation", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}
