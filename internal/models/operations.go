package models

import (
	"context"
	"time"
)

/*
Liest die Operationen aus der Datenbank als View
*/
func (m *DBModel) GetOperations(plantId int) ([]Operation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var ops []Operation
	sql := `
	SELECT 
	id, unit_name, author, plant_name, description
	FROM all_operations
	WHERE plant_id = ?
	`
	rows, err := m.DB.QueryContext(ctx, sql, plantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var op Operation
		err = rows.Scan(
			&op.Id,
			&op.Name,
			&op.Author,
			&op.Plant,
			&op.Description,
		)
		if err != nil {
			return nil, err
		}
		ops = append(ops, op)
	}
	return ops, nil
}
