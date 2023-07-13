package main

import (
	"github.com/schlucht/fhxreader/internal/models"
	"github.com/schlucht/fhxreader/internal/parser"
)

// Speicher eine OP FHX in die Datenbank
func (app *application) insertOperations(txt string) error {

	parsedOP, err := parser.NewFhxString(txt)
	if err != nil {
		return err
	}
	// Durchläuft die geparsten OP zum speichern
	for _, ops := range parsedOP {
		for _, o := range ops.OPs {

			opUnit := models.Unit{
				Name:        o.UnitName,
				Position:    o.UnitPosition,
				Author:      o.Author,
				Description: o.Description,
				Time:        o.Time,
				Type:        1,
			}
			opId, err := app.DB.InsertUnit(opUnit, 1)
			if err != nil {
				return err
			}
			for _, p := range o.Parameters {
				opParams := models.Parameter{
					Name:        p.Name,
					UnitID:      opId,
					Description: p.Description,
				}
				paramId, err := app.DB.InsertParameter(opParams)
				if err != nil {
					return err
				}
				val := models.Value{
					StringValue: p.Value.StringValue,
					ValueSet:    p.Value.Set,
					Hight:       p.Value.High,
					Low:         p.Value.Low,
					CV:          p.Value.Cv,
					ParamId:     paramId,
				}
				_, err = app.DB.InsertValue(val)
				if err != nil {
					return err
				}
			}
		}

	}
	return nil
}