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
	Type        int       `json:"type"`
	Name        string    `json:"name"`
	Position    string    `json:"position"`
	Author      string    `json:"author"`
	Description string    `json:"desc"`
	Time        int       `json:"time"`
	CreatetAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type Parameter struct {
	Id          int       `json:"id"`
	UnitID      int       `json:"unit_id"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	Description string    `json:"desc"`
	CreatetAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type Value struct {
	Id          int       `json:"id"`
	ParamId     int       `json:"param_id"`
	StringValue string    `json:"string_value"`
	ValueSet    string    `json:"value_set"`
	Hight       int       `json:"hight"`
	Low         int       `json:"low"`
	CV          int       `json:"CV"`
	Unit        int       `json:"unit"`
	CreatetAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

// Values in Datenbank speichern
func (m *DBModel) InsertValue(val Value) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		INSERT INTO paramvalues
			(stringvalue, value_set, high, low, cv, unit, created_at, updated_at, param_id)
		Values
			(?,?,?,?,?,?,?,?,?)
	`

	result, err := m.DB.ExecContext(ctx, stmt,
		val.StringValue,
		val.ValueSet,
		val.Hight,
		val.Low,
		val.CV,
		val.Unit,
		time.Now(),
		time.Now(),
		val.ParamId,
	)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rows), nil
}

// Einzelne Unit speichern und die ID zurückgeben
func (m *DBModel) InsertParameter(param Parameter) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		INSERT INTO parameters
			(parameter_name, unit_id, description, created_at, updated_at)
			values(?,?,?,?,?)`
	result, err := m.DB.ExecContext(ctx, stmt,
		param.Name,
		param.UnitID,
		param.Description,
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

// Einzelne Unit speichern und die ID zurückgeben
func (m *DBModel) InsertUnit(u Unit, typ int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		INSERT INTO units
			(unit_name, type_id, position, author, time, description, created_at, updated_at)
			values(?,?,?,?,?,?,?,?)`
	result, err := m.DB.ExecContext(ctx, stmt,
		u.Name,
		typ,
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
