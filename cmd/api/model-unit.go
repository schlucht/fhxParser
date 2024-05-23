package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/schlucht/fhxreader/internal/models"
	"github.com/schlucht/fhxreader/internal/parser"
)

// Der fhx Text durchlaufen und die Unit speichern. Auslesen der Unit,
// extrahieren der Parameter und die Operationen.
//
// Parameters:
//   - fhx: FHX Text
//   - plantID: ID der Anlage
//
// Return:
//   - error: Fehlermeldung
func (app *application) SaveUnit(fhx parser.Fhx, plantID uuid.UUID) error {
	for _, u := range fhx.Units {
		// 1. Unit in die Tabelle speichern
		unitId, _ := app.DB.UnitIdFromName(u.UnitName, plantID)
		if unitId == uuid.Nil {
			unitId = uuid.New()
			var unitModel = models.Unit{
				ID:           unitId,
				PlantID:      plantID,
				UnitName:     u.UnitName,
				UnitCategory: u.UnitCategory,
				UnitPosition: u.UnitPosition,
				Author:       u.Author,
				Description:  u.Description,
				Time:         u.Time,
			}
			err := app.DB.NewUnit(unitModel)
			if err != nil {
				return err
			}
		}
		// 2. Operationen in die Tabelle speichern
		for _, o := range u.Steps {

			op := models.UnitOP{
				ID:            uuid.New(),
				UnitID:        unitId,
				OpKey:         o.Key,
				OpName:        o.Name,
				OpDescription: o.Description,
				OpPosition:    o.Rect,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			}
			uopUnit, _ := app.DB.OpUnitIdFromId(op.OpKey, unitId)
			if uopUnit == uuid.Nil {
				err := app.DB.SaveUnitOps(op)
				if err != nil {
					return err
				}
			}
		}
		// 3. Values zu den Operationen sammeln
		// a) Step durchlaufen
		// b) Step-Parametername speichern
		// c) Kontrolle wenn Origin = CONSTANT dann die Values aus der OP Tabelle holen und
		// und die Werte aus den step_attribute ersetzten
		// d) Wenn der Origin = DEFERRED dann die Values aus der OP Tabelle holen und die Werte durch die Values
		// aus den Paramter Values ersetzen

	}
	return nil
}
