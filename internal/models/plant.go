package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Plant struct {
	ID   string `json:"plant_id"`
	Name string `json:"plant"`
}

func (m *DBModel) CreateNewPlant(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	id := uuid.New()
	stmt := `INSERT INTO plants (plant_id, plant) VALUES(?,?)`
	_, err := m.DB.ExecContext(ctx, stmt, id.String(), name)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) GetPlants() ([]Plant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `SELECT plant_id, plant FROM plants`
	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var plants []Plant
	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		plants = append(plants, Plant{ID: id, Name: name})
	}
	return plants, nil
}
