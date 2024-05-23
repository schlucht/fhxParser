package main

import (
	"github.com/google/uuid"
	"github.com/schlucht/fhxreader/internal/models"
	"github.com/schlucht/fhxreader/internal/parser"
)

func (app *application) SaveOperation(fhx parser.Fhx, plantID uuid.UUID) error {
	for _, o := range fhx.OPs {
		// search for existing operation
		opUUID, err := app.DB.OpFromName(o.UnitName)
		if err != nil {
			app.errorLog.Println("Failed to get operation: ", err)
			return err
		}
		if opUUID == uuid.Nil {
			opUUID = uuid.New()
			var opModel = models.Operation{
				ID:     opUUID,
				OPName: o.UnitName,
			}
			err = app.DB.NewOP(opModel)
			if err != nil {
				app.errorLog.Println("Failed to create new operation: ", err)
				return err
			}
		}

		// search exist OperationPlant
		existPlant, err := app.DB.OPPlantFromID(opUUID, plantID)
		if err != nil {
			app.errorLog.Println("Failed to get operation plant: ", err)
			return err
		}
		if existPlant == uuid.Nil {
			existPlant := uuid.New()
			var procedureModel = models.OperationPlant{
				ID:          existPlant,
				PlantID:     plantID,
				OperationID: opUUID,
				Category:    o.UnitCategory,
				Position:    o.UnitPosition,
				Author:      o.Author,
				Description: o.Description,
				OPTime:      o.Time,
			}
			err = app.DB.NewOPPlant(procedureModel)
			if err != nil {
				app.errorLog.Println("Failed to create new procedure: ", err)
				return err
			}
		}

		for _, p := range o.Parameters {

			var parameterModel = models.Parameter{
				ID:          uuid.New(),
				OPPlantID:   existPlant,
				Name:        p.Name,
				Description: p.Description,
			}

			// search exist Parameter
			existParam, err := app.DB.ExistParam(existPlant, p.Name)
			if err != nil {
				app.errorLog.Println("Failed to get parameter: ", err)
				return err
			}
			if existParam != uuid.Nil {
				err = app.DB.NewParam(parameterModel, existPlant)
				if err != nil {
					app.errorLog.Println("Failed to create new parameter: ", err)
					return err
				}
			}
			app.infoLog.Println("existParam: ", parameterModel.ID.String())
			var valueModel = models.Value{
				ID:          uuid.New(),
				ParamID:     parameterModel.ID,
				High:        p.Value.High,
				Low:         p.Value.Low,
				Cv:          p.Value.Cv,
				Unit:        p.Value.Unit,
				Set:         p.Value.Set,
				StringValue: p.Value.StringValue,
			}
			err = app.DB.NewValue(valueModel)
			if err != nil {
				app.errorLog.Println("Failed to create new value: ", err)
				return err
			}
		}
	}

	return nil
}
