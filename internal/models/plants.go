package models

import (
	"context"
	"time"
)

// Alle Plants(Anlagen) aus der Datenbank auslesen
func (m *DBModel) LoadAllPlants() ([]*Plant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var plants []*Plant
	query := `SELECT id, plant_name, created_at, updated_at FROM plants`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Plant

		err = rows.Scan(
			&p.Id,
			&p.Plant,
			&p.CreatetAt,
			&p.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		plants = append(plants, &p)
	}
	return plants, nil
}

// Löscht eine Anlage aus der Datenbank
func (m *DBModel) DeletePlant(id int) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM plants WHERE id = ?`

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Anlage Name ändern
func (m *DBModel) UpdatePlant(plant Plant) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE plants SET
				plant_name = ?,			
				updated_at = ?
				WHERE id = ?

			`
	result, err := m.DB.ExecContext(ctx, query, plant.Plant, time.Now(), plant.Id)
	if err != nil {
		return 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

