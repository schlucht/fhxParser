package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/schlucht/fhxreader/internal/models"
)

type Recipes struct {
	Recipes []models.Recipe
	Count   int
}

func (app *application) insertRecipe(parsedFhx []models.Recipe, plantID int) error {

	// 2. Durchlaufen der Steps
	// 3. Speichern der Steps Namen in den Units
	recipes := parsedFhx
	plant := plantID

	for _, recipe := range recipes {
		// 1. Speichern des Rezeptnamen
		exist, err := app.DB.ExistRecipe(recipe.Name)
		if err != nil {
			return err
		}
		if !exist {
			app.DB.InsertRecipe(recipe, plant)
		} else {
			fmt.Println("Rezept existiert bereits")
		}
	}
	return nil
}

func (app *application) getAllRecipes(plantId int) (Recipes, error) {
	res, err := app.DB.GetAllRecipes(plantId)
	recipes := Recipes{Recipes: res, Count: len(res)}
	if err != nil {
		app.errorLog.Println(err)
		return recipes, err
	}
	return recipes, nil
}
