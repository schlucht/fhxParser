package models

import (
	"context"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

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
func (m *DBModel) InsertUnit(u Unit, typ int, plant_id int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		INSERT INTO units
			(unit_name, type_id, plant_id, position, author, time, description, created_at, updated_at)
			values(?,?,?,?,?,?,?,?,?)`
	result, err := m.DB.ExecContext(ctx, stmt,
		u.Name,
		typ,
		plant_id,
		u.Position,
		u.Author,
		u.Time,
		u.Description,
		time.Now(),
		time.Now(),
	)

	log.Println(err.(*mysql.MySQLError).Number)
	if err != nil {
		m.DBError = int(err.(*mysql.MySQLError).Number)
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
