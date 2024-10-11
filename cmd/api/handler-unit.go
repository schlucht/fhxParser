package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Alle Units auf von der Anlage auf die Seite
// laden
func (app *application) UnitPage(w http.ResponseWriter, r *http.Request) {
	plantId := chi.URLParam(r, "plantId")

	uuid, err := uuid.Parse(plantId)
	if err != nil {
		app.badRequest(w, err, "UUID nicht IO", http.StatusInternalServerError)
		return
	}

	unit, err := app.DB.UnitIdFromPlantId(uuid)
	if err != nil {
		app.badRequest(w, err, "GetPlantsUnit", http.StatusNoContent)
		return
	}

	// Daten an das Frontend übergeben
	data := make(map[string]interface{})
	data["units"] = unit

	if err := app.renderTemplate(w, r, "unit", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}
