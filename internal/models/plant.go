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

// Tabelle mit 3 Betrieben erstellen.
func (m *DBModel) InsertNewPlants() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.createPlantTable()
	if err != nil {
		return err
	}

	// Kontrolle ob Inhalt vorhanden
	rows, err := m.DB.QueryContext(ctx, "SELECT COUNT(*) FROM plants")
	if err != nil {
		return err
	}
	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return err
		}
	}

	// Wenn keine Daten in DB drei Anlagen in die DB speichern
	if count == 0 {
		plants := []string{"leer"}
		for _, p := range plants {
			err := m.CreateNewPlant(p)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Neuer Betrieb speichern und in die Datenbank speichern.
func (m *DBModel) CreateNewPlant(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	id := uuid.New()
	stmt := `INSERT INTO plants (plant_id, plant, updated_at, created_at) VALUES(?,?,?,?)`
	_, err := m.DB.ExecContext(ctx, stmt, id.String(), name, time.Now(), time.Now())
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
	err = m.createPlantTable()
	if err != nil {
		return nil, err
	}

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
