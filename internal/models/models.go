package models

import (
	"context"
	"database/sql"
	"time"
)

// for Database connection
type DBModel struct {
	DB *sql.DB
}

// Wrapper for all Models
type Models struct {
	DB DBModel
}

// New Models return a Model
func NewModel(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Unit struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Position    string    `json:"position"`
	Author      string    `json:"author"`
	Description string    `json:"desc"`
	Time        int       `json:"time"`
	CreatetAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

// Speichern aller Units Return Letzte ID
func (m *DBModel) InsertUnit(u Unit) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		INSERT INTO units
			(unit_name, position, author, time, description, created_at, updated_at)
			values(?,?,?,?,?,?,?)`
	result, err := m.DB.ExecContext(ctx, stmt,
		u.Name,
		u.Position,
		u.Author,
		u.Time,
		u.Description,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Auslesen aller Units return []UNIT, ERROR
func (m *DBModel) GetUnits() ([]Unit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var units []Unit
	sql := `
		SELECT 
			id, unit_name, position, author, time, description, created_at, updated_at FROM units
	`
	rows, err := m.DB.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var unit Unit
		err = rows.Scan(
			&unit.Id,
			&unit.Name,
			&unit.Position,
			&unit.Author,
			&unit.Time,
			&unit.Description,
			&unit.CreatetAt,
			&unit.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		units = append(units, unit)
	}
	return units, nil

}
