package main

import (
	"strings"
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
		// 1. Unit in die Tabelle speichern oder aktualisieren
		unitId, err := app.saveUnit(u, plantID)
		if err != nil {
			app.errorLog.Println("Failed to save new unit: ", err)
			return err
		}

		for _, o := range u.Steps {

			// 2. Operationen in die Tabelle speichern oder aktualisieren
			err := app.saveUnitOP(unitId, o)
			if err != nil {
				app.errorLog.Println("Failed to save new operation: ", err)
				return err
			}

			// 3. Values zusammensuchen und Speichern
			for _, p := range o.StepParameters {
				app.infoLog.Println(p)
				// val := app.getValues(o.Key, p.Name, u, o, plantID)
				// err := app.DB.SaveUnitValue(val)
				// if err != nil {
				// 	return err
				// }
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

func (app *application) saveUnitOP(unitId uuid.UUID, o parser.Step) error {
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
	return nil
}

// Speichert die Unit in der Datenbank
// Ist die Unit bereits in der Datenbank vorhanden wird der DS aktualisiert
// Parameter:
//   - parser.Unit parserUnit struct
//   - uuid.UUID plantID
//
// Return:
//   - uuid.UUID: ID der Unit
//   - error: Fehlermeldung
func (app *application) saveUnit(u parser.Unit, plantID uuid.UUID) (uuid.UUID, error) {
	unitId, _ := app.DB.UnitIdFromName(u.UnitName, plantID)

	var unitModel = models.Unit{
		PlantID:      plantID,
		UnitName:     u.UnitName,
		UnitCategory: u.UnitCategory,
		UnitPosition: u.UnitPosition,
		Author:       u.Author,
		Description:  u.Description,
		Time:         u.Time,
	}

	if unitId == uuid.Nil {
		unitId = uuid.New()
		unitModel.ID = unitId
		err := app.DB.SaveUnit(unitModel)
		if err != nil {
			app.errorLog.Println("Failed to create new unit: ", err)
			return uuid.Nil, err
		}
	} else {
		unitModel.ID = unitId
		err := app.DB.UpdateUnit(unitModel)
		if err != nil {
			app.errorLog.Println("Failed to update unit: ", err)
			return uuid.Nil, err
		}
	}
	return unitId, nil
}

func (app *application) getValues(opKey string, paramName string, u parser.Unit, s parser.Step, plant uuid.UUID) models.UnitValue {
	val := models.UnitValue{}
	stepParams := s.StepParameters

	for _, stepParam := range stepParams {
		isConstan := strings.Index(stepParam.Origin, "CONSTANT")
		app.infoLog.Println(len(stepParams))
		if isConstan > -1 {
			// Suchen in den OP Tabelle nache dem OP Namen und nach der Anlage
			// Mit der gefunden ID in der Parameter Tabelle Value suchen und die neusten Values auslesen
			opId, err := app.DB.IDOPPlantFromName(opKey, plant)
			if err != nil {
				return val
			}
			app.infoLog.Println(opId)
		}

		// val := models.UnitValue{
		// 	ID:          uuid.New(),
		// 	UnitID:      unitId,
		// 	StringValue: p.StringValue,
		// 	Set:         p.Value.Set,
		// 	High:        p.Value.High,
		// 	Low:         p.Value.Low,
		// 	Cv:          p.Value.CV,
		// 	Unit:        p.Value.Unit,
		// 	CreatedAt:   time.Now(),
		// 	UpdatedAt:   time.Now(),
		// }
	}
	return val
}
