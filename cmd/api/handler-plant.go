package main

import (
	"net/http"
)

// Öffnet die Plant Webseite und übergibt Daten
func (app *application) PlantPage(w http.ResponseWriter, r *http.Request) {

	// Alle Anlagen aus der Datenbank auslesen
	plants, err := app.DB.GetPlants()
	if err != nil {
		app.badRequest(w, r, err, "GetPlants")
		return
	}

	// Daten an das Frontend übergeben
	data := make(map[string]interface{})
	data["plants"] = plants
	data["counts"] = len(plants)

	if err := app.renderTemplate(w, r, "plant", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}

// Alle Betriebe aus der Datenbank auslesen und als
// JSON übergeben
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

// Speicher einer Anlage in einer
// Datenbank
func (app *application) PlantSave(w http.ResponseWriter, r *http.Request) {
	var plantInput struct {
		Name string `json:"plantName"`
	}

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

// Loeschen einer Anlage aus der Datenbank
func (app *application) PlantDelete(w http.ResponseWriter, r *http.Request) {
	var plantId struct {
		ID string `json:"plantId"`
	}
	err := app.readJSON(w, r, &plantId)
	if err != nil {
		app.badRequest(w, r, err, "PlantDelete: Read JSON")
		return
	}
	err = app.DB.PlantDelete(plantId.ID)
	if err != nil {
		app.badRequest(w, r, err, "PlantDelete: DeletePlant")
		return
	}
	w.WriteHeader(http.StatusAccepted)
	http.Redirect(w, r, "/plants", http.StatusSeeOther)
}

// Anlage aktualisieren
func (app *application) PlantUpdate(w http.ResponseWriter, r *http.Request) {
	var plantInput struct {
		ID   string `json:"plantId"`
		Name string `json:"plantName"`
	}
	err := app.readJSON(w, r, &plantInput)
	if err != nil {
		app.badRequest(w, r, err, "PlantUpdate: Read JSON")
		return
	}
	err = app.DB.PlantUpdate(plantInput.Name)
	if err != nil {
		app.badRequest(w, r, err, "PlantUpdate: UpdatePlant")
		return
	}
	w.WriteHeader(http.StatusAccepted)
	http.Redirect(w, r, "/plants", http.StatusSeeOther)
}
