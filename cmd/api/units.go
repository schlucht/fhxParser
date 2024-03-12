package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/schlucht/fhxreader/internal/models"
	"github.com/schlucht/fhxreader/internal/parser"
)

type Units struct {
	Units []models.Unit
	Count int
}

// Speichert eine OP FHX in die Datenbank
func (app *application) insertUnit(parsedUP parser.Fhx, plantID int) ([]string, error) {

	var errorString = []string{}

	// Durchläuft die geparsten OP zum speichern
	for _, o := range parsedUP.Units {

		opUnit := models.Unit{
			Name:        o.UnitName,
			Position:    o.UnitPosition,
			Author:      o.Author,
			Description: o.Description,
			Time:        o.Time,
			Type:        o.Type,
		}

		opId, err := app.DB.InsertUnit(opUnit, 2, plantID)

		if err != nil {
			nb, ok := err.(*mysql.MySQLError)
			if !ok {
				return nil, err
			}
			if nb.Number == 1062 {
				errorString = append(errorString, o.UnitName)
				continue //return errors.New("duplicate unit")
			}
		}

		for _, p := range o.Parameters {

			opParams := models.Parameter{
				Name:        p.Name,
				UnitID:      opId,
				Description: p.Description,
			}
			paramId, err := app.DB.InsertParameter(opParams)
			if err != nil {
				log.Println(err)
				return nil, err
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
				return nil, err
			}
		}
	}
	return errorString, nil
}

func (app *application) AllUnits(plantId int) (Units, error) {

	list := Units{}
	ops, err := app.DB.GetUnits(plantId)

	if err != nil {
		return list, err
	}
	list.Units = ops
	list.Count = len(ops)

	return list, nil
}
