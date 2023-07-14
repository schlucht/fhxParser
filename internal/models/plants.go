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
