package main

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/schlucht/fhxreader/internal/models"
	"github.com/schlucht/fhxreader/internal/parser"
)

type Operations struct {
	Operations []models.Operation
	Count      int
}

// Speichert eine OP FHX in die Datenbank
func (app *application) insertOperations(parsedOP parser.Fhx, plantID int) ([]string, error) {

	var errorString = []string{}

	// Durchläuft die geparsten OP zum speichern
	for _, o := range parsedOP.OPs {

		opUnit := models.Unit{
			Name:        o.UnitName,
			Position:    o.UnitPosition,
			Author:      o.Author,
			Description: o.Description,
			Time:        o.Time,
			Type:        o.Type,
		}
		opId, err := app.DB.InsertUnit(opUnit, 1, plantID)
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


/*
Alle Operationen aus der DB dem Handler zur Verfügung stellen
*/
func (app *application) AllOperations(plantId int) (Operations, error) {

	list := Operations{}
	ops, err := app.DB.GetOperations(plantId)

	if err != nil {
		return list, err
	}
	list.Operations = ops
	list.Count = len(ops)

	return list, nil
}

/*
Die Parameter einer Operation aus der Datenbank auslesen
*/
func (app *application) GetOpFromId(opId int) (models.Operation, error) {

	op, err := app.DB.GetOperationFromId(opId)
	if err != nil {
		return op, err
	}
	return op, nil
}

// Operation in der Datenbank speichern
func (app *application) saveOperations(fhx parser.Fhx, plantId int) (string, error) {
	doubleUp, err := app.insertOperations(fhx, plantId)
	if err != nil {
		return "", err
	}
	if len(doubleUp) > 0 {
		s := strings.Join(doubleUp, ",")
		return fmt.Sprintf("Dopplete UP's: %s", s), nil
	}
	return "", nil
}

