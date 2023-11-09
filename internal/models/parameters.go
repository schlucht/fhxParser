package models

import (
	"context"
	"time"
)

/* Die Parameter einer Operation anhand der OP ID */
func (m *DBModel) GetParamsFromOpId(opId int) ([]Parameter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var params []Parameter

	query := `
		SELECT 
		param_id, parameter_name, created_at, updated_at, description, unit_id
		FROM parameters
		WHERE unit_id = ?		
	`
	rows, err := m.DB.QueryContext(ctx, query, opId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var param Parameter
		err = rows.Scan(
			&param.Id,
			&param.Name,
			&param.CreatetAt,
			&param.UpdatedAt,
			&param.Description,
			&param.UnitID,
		)
		if err != nil {
			return nil, err
		}
		v, err := m.GetValueFromId(param.Id)
		if err != nil {
			return nil, err
		}
		param.Value = v
		params = append(params, param)
	}

	return params, nil
}

// Auslesen der Values zu einem Parameter
func (m *DBModel) GetValueFromId(paramId int) (Value, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var v Value
	query := `
		SELECT 
		value_id, stringvalue, value_set, high, low, cv, unit, created_at, updated_at, param_id
		FROM paramvalues
		WHERE param_id = ?		
	`
	err := m.DB.QueryRowContext(ctx, query, paramId).Scan(
		&v.Id,
		&v.StringValue,
		&v.ValueSet,
		&v.Hight,
		&v.Low,
		&v.CV,
		&v.Unit,
		&v.CreatetAt,
		&v.UpdatedAt,
		&v.ParamId,
	)
	return v, err
}
