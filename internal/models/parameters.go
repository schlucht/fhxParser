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
		params = append(params, param)
	}
	return params, nil
}
