package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// Alle Rezepte einlesen laut Anlage
func (app *application) getAllRecipesHandler(w http.ResponseWriter, r *http.Request) {

	f, err := io.ReadAll(r.Body)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Keine Daten vorhanden")
		app.badRequest(w, r, err)
		return
	}
	var re antwort

	err = json.Unmarshal(f, &re)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Betrieb fehlt konnte nicht geparst werden")
		app.badRequest(w, r, err)
		return
	}

	recipes, err := app.DB.GetAllRecipes(re.Id)
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, r, err)
		return
	}
	rec := Recipes{
		Recipes: recipes,
		Count:   len(recipes),
	}
	res, err := json.Marshal(recipes)
	if err != nil {
		app.errorLog.Println("JSON Recipe nicht funktioniert", err)
		app.badRequest(w, r, err)
		return
	}
	if rec.Count == 0 {
		j.OK = false
		j.Content = "{}"
		j.Message = "Kein Rezept gefunden"
	} else {
		j.OK = true
		j.Message = "Daten OK"
		j.Content = string(res)
	}

	//app.infoLog.Println("Get all recipes")
	w.Header().Set("Content-Type", "application/json")
	app.writeJSON(w, http.StatusOK, recipes)
}
