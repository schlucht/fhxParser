package main

import (
	"net/http"
)

func (app *application) PlantPage(w http.ResponseWriter, r *http.Request) {

	plants, err := app.DB.GetPlants()
	if err != nil {
		app.badRequest(w, r, err, "GetPlants")
		return
	}
	data := make(map[string]interface{})
	data["plants"] = plants

	if err := app.renderTemplate(w, r, "plant", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) AllPlants(w http.ResponseWriter, r *http.Request) {
	plants, err := app.DB.GetPlants()
	if err != nil {
		app.badRequest(w, r, err, "GetPlants")
		return
	}
	if err = app.writeJSON(w, http.StatusOK, plants); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) PlantSave(w http.ResponseWriter, r *http.Request) {
	var plantInput struct {
		Name string `json:"plantName"`
	}
	app.infoLog.Println(plantInput.Name)

	err := app.readJSON(w, r, &plantInput)
	if err != nil {
		app.badRequest(w, r, err, "PlantSave: readJson")
		return
	}
	err = app.DB.CreateNewPlant(plantInput.Name)
	if err != nil {
		app.badRequest(w, r, err, "PlantSave: CreateNewPlant")
		return
	}

	w.WriteHeader(http.StatusAccepted)
	http.Redirect(w, r, "/plants", http.StatusSeeOther)
}
