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

				err = app.getValues(o.Key, o.Name, p, plantID, unitId)
				if err != nil {
					app.errorLog.Println("Failed to get values: ", err)
					return err
				}

				if strings.Contains(p.Origin, "DEFERRED") {
					for _, up := range u.Parameters {
						if up.Name == p.DeferTo {
							// hier die neuen Values in DB speichern
							unitops, err := app.DB.OpUnitIdFromName(o.Name, unitId)
							// app.infoLog.Println(unitops, unitId, o.Name)
							if err != nil {
								app.errorLog.Println("Failed to get unit op  id: ", err)
								return err
							}
							id, _ := app.DB.IDUnitParamDeferTo(unitops, p.DeferTo)
							app.infoLog.Println(p.DeferTo, id)
							if id == uuid.Nil {
								app.errorLog.Println("Failed to get param id: ", id, p.DeferTo)
								// return errors.New("now id in db")
							}
							var val = models.UnitValue{
								ID:          uuid.New(),
								UnitID:      id,
								High:        up.Value.High,
								Low:         up.Value.Low,
								Cv:          up.Value.Cv,
								Unit:        up.Value.Unit,
								Set:         up.Value.Set,
								StringValue: up.Value.StringValue,
							}
							err = app.DB.SaveUnitParamValue(val)
							if err != nil {
								app.errorLog.Println("Failed to update value: ", err)
								return err
							}
						}
					}
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
	uopUnit, _ := app.DB.OpUnitIdFromName(op.OpName, unitId)
	if uopUnit == uuid.Nil {
		err := app.DB.SaveUnitOps(op)
		if err != nil {
			app.errorLog.Println("Failed to save new operation: ", err)
			return err
		}
	} else {
		// app.infoLog.Println("updateUnitOP else teil: ", op.OpKey, uopUnit)
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

func (app *application) getValues(
	opKey string,
	opName string,
	p parser.StepParameter,
	plant uuid.UUID,
	unitid uuid.UUID) error {

	opId, err := app.DB.IDOPPlantFromName(opKey, plant)
	if err != nil {
		app.errorLog.Println("IDOPPlantFromName", err)
		return err
	}

	paramId, err := app.DB.ParamIdFromName(p.Name, opId)
	if err != nil {
		app.errorLog.Println("ParamIdFromName", err)
		return err
	}

	val, err := app.DB.LastValueFromParamId(paramId)
	if err != nil {
		app.errorLog.Println("LastValueFromParamId", err)
		return err
	}

	unitOP, err := app.DB.OpUnitIdFromName(opName, unitid)
	
	if err != nil {
		app.errorLog.Println("OpUnitIdFromName", err)
		return err
	}
	var param models.UnitParameter
	param.ID = uuid.New()
	param.UnitOPID = unitOP
	param.ParamName = p.Name
	param.OriginValueID = val.ID
	param.Origin = p.Origin
	param.Deferto = p.DeferTo
	parId, _ := app.DB.ExistUnitParam(param.ID, param.ParamName)

	if parId == uuid.Nil {
		app.infoLog.Println(opId, p.Name)
		err = app.DB.SaveUnitParameter(param)
		if err != nil {
			app.errorLog.Println("SaveUnitParemter", err)
			return err
		}
	}

	return nil
}
