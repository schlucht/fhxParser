package main

import "github.com/schlucht/fhxreader/internal/models"

type OPParameters struct {
	Unit      models.Operation
	ParamList []Parameter
}

type Parameter struct {
	Name        string       `json:"name"`
	Description string       `json:"desc"`
	Value       models.Value `json:"value"`
}

/*
Liest die Parameter laut unitId aus
*/
func (app *application) GetParamFromOPId(opId int) (OPParameters, error) {
	list := OPParameters{}

	op, err := app.DB.GetOperationFromId(opId)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Probleme beim Einesen einer Unit")
		return list, err
	}

	ps, err := app.DB.GetParamsFromOpId(opId)
	if err != nil {
		app.errorLog.Printf("%v, %s", err, "Parameter konnte nicht gelesen werden")
		return list, err
	}
	var params []Parameter
	var par Parameter
	for _, p := range ps {
		par = Parameter{
			Name:        p.Name,
			Description: p.Description,
			Value: p.Value,
		}
		params = append(params, par)
	}
	list.ParamList = params
	list.Unit = op

	return list, nil
}
