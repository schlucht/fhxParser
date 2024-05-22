package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Plant struct {
	ID        uuid.UUID `json:"plant_id"`
	Name      string    `json:"plant"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
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
	stmt := `SELECT plant_id, plant, created_at, updated_at FROM plants`
	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var plants []Plant
	for rows.Next() {
		var p Plant

		err := rows.Scan(&p.ID, &p.Name, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		plants = append(plants, p)
	}
	return plants, nil
}

func (m *DBModel) PlantSave(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `UPDATE plants SET plant = ?, updated_at = ?`
	_, err := m.DB.ExecContext(ctx, stmt, name, time.Now())
	if err != nil {
		return err
	}
	return nil
}
