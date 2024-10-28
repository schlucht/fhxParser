package main

import (
	"github.com/google/uuid"
	"github.com/schlucht/fhxreader/internal/models"
	"github.com/schlucht/fhxreader/internal/parser"
)

func (app *application) SaveOperation(fhx parser.Fhx, plantID uuid.UUID) error {
	err := app.DB.CreateOPTable()
	if err != nil {
		return err
	}
	for _, o := range fhx.OPs {
		// OP Model speichern
		opModel, err := app.saveOP(o)
		if err != nil {
			app.errorLog.Println("Save saveOP: ", err)
			return err
		}

		// OP anhand der Anlage speichern
		opPlantID, err := app.saveOPPlant(opModel, plantID, o)
		if err != nil {
			app.errorLog.Println("Save saveOPPlant: ", err)
			return err
		}

		for _, p := range o.Parameters {

			// Parameter speichern
			paramID, err := app.saveParams(opPlantID, p)
			if err != nil {
				app.errorLog.Println("Failed to create new parameter: ", err)
				return err
			}
			// Values werden immer hinten Angehängt bei jedem Upload
			err = app.saveValue(paramID, p)
			if err != nil {
				app.errorLog.Println("Failed to create new value: ", err)
				return err
			}
		}
	}

	return nil
}

// Speichert die Operation in der Datenbank
// Ist die Operation bereits in der Datenbank vorhanden wird der DS aktualisiert
// Parameter: parser.Unit parserUnit struct
// Return:
//   - models.Operation struct,
//   - error
func (app *application) saveOP(parserUnit parser.Unit) (models.Operation, error) {
	var opModel = models.Operation{
		OPName: parserUnit.UnitName,
	}

	op, err := app.DB.OpFromName(parserUnit.UnitName)
	if err != nil {
		app.errorLog.Println("Failed to get operation: ", err)
		return opModel, err
	}

	opModel = op
	if op.IsEmpty() {
		opModel.ID = uuid.New()
		opModel.OPName = parserUnit.UnitName
		err = app.DB.NewOP(opModel)
		if err != nil {
			app.errorLog.Println("Failed to create new operation: ", err)
			return opModel, err
		}
	} else {
		// TODO: Update Operation
		opModel.ID = op.ID
		err = app.DB.UpdateOP(opModel)
		if err != nil {
			app.errorLog.Println("Failed to update operation: ", err)
			return opModel, err
		}
	}
	return opModel, nil
}

// Kontrolliert schon eine OP mit einer Anlage zusamme gespeichert ist
// Wenn nich wird eine neue Angelegt
// Wenn vorhanden wird der DS aktualisiert
// Parameter:
// 	- models.Operation struct
// 	- uuid.UUID plantID
// 	- parser.Unit parserUnit struct

// Return:
//   - PlantID UUID,
//   - error
func (app *application) saveOPPlant(opModel models.Operation, plantID uuid.UUID, parserUnit parser.Unit) (uuid.UUID, error) {
	// sucht ob die Operation für die Anlage  bereits in der Datenbank existiert
	existPlant, err := app.DB.OPPlantFromID(opModel.ID, plantID)
	if err != nil {
		app.errorLog.Println("Failed to get operation plant: ", err)
		return uuid.Nil, err
	}

	if existPlant == uuid.Nil {
		existPlant = uuid.New()
		var procedureModel = models.OperationPlant{
			ID:          existPlant,
			PlantID:     plantID,
			OperationID: opModel.ID,
			Category:    parserUnit.UnitCategory,
			Position:    parserUnit.UnitPosition,
			Author:      parserUnit.Author,
			Description: parserUnit.Description,
			OPTime:      parserUnit.Time,
		}
		err = app.DB.NewOPPlant(procedureModel)
		if err != nil {
			app.errorLog.Println("Failed to create new procedure: ", err)
			return uuid.Nil, err
		}
	}
	return existPlant, nil
}

// Speichert die Parameter in der Datenbank
// Wenn die Parameter bereits in der Datenbank existieren wird der DS aktualisiert
// Parameter:
//   - parser.Unit parserUnit struct
//   - uuid.UUID plantID
//
// Return:
//   - error
func (app *application) saveParams(opPlantID uuid.UUID, p parser.Parameter) (uuid.UUID, error) {

	var parameterModel = models.Parameter{
		OPPlantID:   opPlantID,
		Name:        p.Name,
		Description: p.Description,
	}

	// Existierender Parameter mit suchen
	existParam, err := app.DB.ExistParam(opPlantID, p.Name)
	if err != nil {
		app.errorLog.Println("Failed to get parameter: ", err)
		return uuid.Nil, err
	}

	// Wenn keine Id gefunden wurde dann Speichern
	// Wenn eine Id gefunden wurde dann Udate
	if existParam == uuid.Nil {
		existParam = uuid.New()
		parameterModel.ID = existParam
		err = app.DB.NewParam(parameterModel, opPlantID)
		if err != nil {
			app.errorLog.Println("Failed to create new parameter: ", err)
			return uuid.Nil, err
		}
	} else {

		parameterModel.ID = existParam
		err = app.DB.UpdateParams(parameterModel, opPlantID)
		if err != nil {
			app.errorLog.Println("Failed to update parameter: ", err)
			return uuid.Nil, err
		}
	}
	return existParam, nil
}

// Speichert die Values zu einem Parameter
// Parameter:
//   - uuid.UUID paramID
//   - parser.Parameter p
//
// Return:
//   - error
func (app *application) saveValue(paramID uuid.UUID, p parser.Parameter) error {
	var valueModel = models.Value{
		ID:          uuid.New(),
		ParamID:     paramID,
		High:        p.Value.High,
		Low:         p.Value.Low,
		Cv:          p.Value.Cv,
		Unit:        p.Value.Unit,
		ValueSet:    p.Value.Set,
		StringValue: p.Value.StringValue,
	}
	err := app.DB.NewValue(valueModel)
	if err != nil {
		app.errorLog.Println("Failed to create new value: ", err)
		return err
	}
	return nil
}

func (app *application) OpDetailsFromOp(id string) ([]models.Parameter, error) {
	oppantId := uuid.MustParse(id)
	var params []models.Parameter = []models.Parameter{}

	params, err := app.DB.ParamFromOPPlantID(oppantId)
	if err != nil {
		app.errorLog.Println("Failed to create new value: ", err)
		return params, err
	}
	return params, nil
}
