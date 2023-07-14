package main

import (
	"github.com/schlucht/fhxreader/internal/models"
)

func (app *application) loadPlants() ([]*models.Plant, error) {
	plants, err := app.DB.LoadAllPlants()
	if err != nil {
		app.errorLog.Println(err)
	}
	return plants, nil
}
