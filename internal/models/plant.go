package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Plant struct {
	ID        uuid.UUID `json:"plant_id"`
	Name      string    `json:"plant"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (m *DBModel) CreatePlantTable() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `
		CREATE TABLE IF NOT EXISTS plants (
			plant_id VARCHAR(255) NOT NULL PRIMARY KEY,
			plant VARCHAR(255) NOT NULL,
			updated_at DATETIME NOT NULL,
			created_at DATETIME NOT NULL,
			CONSTRAINT plants_unique UNIQUE (plant)
		);
	`
	_, err := m.DB.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}
	return nil
}

// Neuer Betrieb speichern und in die Datenbank speichern.
func (m *DBModel) CreateNewPlant(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	p, err := m.GetPlantFromName(name)
	if err != nil {
		return err
	}
	if len(p) > 0 {
		return nil
	}
	id := uuid.New()
	stmt := `INSERT INTO plants (plant_id, plant, updated_at, created_at) VALUES(?,?,?,?)`
	_, err = m.DB.ExecContext(ctx, stmt, id.String(), name, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

// Alle Anlagen heraussuchen
func (m *DBModel) GetPlants() ([]Plant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var stmt string
	var rows *sql.Rows
	var err error

	stmt = `SELECT plant_id, plant, created_at, updated_at 
		FROM plants`
	rows, err = m.DB.QueryContext(ctx, stmt)
	if err != nil {
		m.errorLog.Println(err)
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

// Betrieb anhand der ID heraussuchen
func (m *DBModel) GetPlantFromID(uuid string) ([]Plant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var stmt string
	var rows *sql.Rows
	var err error
	stmt = `SELECT plant_id, plant, created_at, updated_at 
	FROM plants WHERE plant_id = ?`
	rows, err = m.DB.QueryContext(ctx, stmt, uuid)
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

func (m *DBModel) GetPlantFromName(name string) ([]Plant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var stmt string
	var rows *sql.Rows
	var err error
	stmt = `SELECT plant_id, plant, created_at, updated_at 
	FROM plants WHERE plant = ?`
	rows, err = m.DB.QueryContext(ctx, stmt, name)
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

func (m *DBModel) PlantUpdate(name string, uuid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `UPDATE plants SET plant = ?, updated_at = ? WHERE plant_id = ?`
	_, err := m.DB.ExecContext(ctx, stmt, name, time.Now(), uuid)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) PlantDelete(uuid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `DELETE FROM plants WHERE plant_id = ?`
	_, err := m.DB.ExecContext(ctx, stmt, uuid)
	if err != nil {
		return err
	}
	return nil
}
