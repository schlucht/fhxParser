package models

import (
	"context"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Unit struct {
	ID           uuid.UUID `json:"id"`
	PlantID      uuid.UUID `json:"plant_id"`
	UnitName     string    `json:"unitname"`
	UnitCategory string    `json:"unitcategory"`
	UnitPosition string    `json:"unitpositon"`
	Author       string    `json:"author"`
	Description  string    `json:"description"`
	Time         int       `json:"time"`
}

type UnitParameter struct {
	ID        uuid.UUID `json:"unitparam_id"`
	UnitOPID  uuid.UUID `json:"unitop_id"`
	ParamName string    `json:"param_name"`
	Deferto   string    `json:"deferto"`
	Origin    string    `json:"origin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UnitOP struct {
	ID            uuid.UUID `json:"unitop_id"`
	UnitID        uuid.UUID `json:"unit_id"`
	OpKey         string    `json:"op_key"`
	OpName        string    `json:"op_name"`
	OpDescription string    `json:"op_descr"`
	OpPosition    string    `json:"op_pos"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UnitValue struct {
	ID          uuid.UUID `json:"value_id"`
	UnitID      uuid.UUID `json:"params_id"`
	StringValue string    `json:"stringvalue,omitempty"`
	Set         string    `json:"value_set,omitempty"`
	High        int       `json:"high,omitempty"`
	Low         int       `json:"low,omitempty"`
	Cv          int       `json:"cv,omitempty"`
	Unit        string    `json:"unit,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Neue Unit in die Tabelle spiechern
// Parameters:
//   - up: Unit Der Eintrag der gespeichert werden soll
//
// Return:
//   - error: Fehlermeldung
func (m *DBModel) NewUnit(up Unit) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if up.ID == uuid.Nil {
		up.ID = uuid.New()
	}

	stmt := `INSERT INTO units (unit_id, plant_id, unit_name, unit_category, unit_pos, unit_time, unit_author, unit_descr) VALUES(?,?,?,?,?,?,?,?)`
	_, err := m.DB.ExecContext(ctx, stmt,
		up.ID,
		up.PlantID,
		up.UnitName,
		up.UnitCategory,
		up.UnitPosition,
		up.Time,
		up.Author,
		up.Description,
	)
	if err != nil {
		me, ok := err.(*mysql.MySQLError)
		if !ok {
			return err
		}
		// Eintrag in DB schon vorhanden
		if me.Number == 1062 {
			return nil
		}
		return err
	}
	return nil
}

// Auslesen einer gespeicherten Unit anhand der Anlage und des Unitnamen
//
// Parameters:
//   - upName: Name der Unit
//   - plantID: ID der Anlage
//
// Return:
//   - uuid.UUID: ID der Unit
//   - error: Fehlermeldung
func (m *DBModel) UnitIdFromName(upName string, plantID uuid.UUID) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT unit_id FROM units WHERE plant_id = ? AND unit_name = ?`
	var id uuid.UUID
	row := m.DB.QueryRowContext(ctx, stmt, plantID, upName)
	err := row.Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (m *DBModel) OpUnitIdFromId(opKey string, unitID uuid.UUID) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT unit_id FROM unit_ops WHERE unit_id = ? AND op_key = ?`

	var id uuid.UUID
	row := m.DB.QueryRowContext(ctx, stmt, unitID, opKey)
	err := row.Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (m *DBModel) SaveUnitOps(up UnitOP) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if up.ID == uuid.Nil {
		up.ID = uuid.New()
	}
	m.errorLog.Printf("UnitOP: %s", up.OpKey)
	stmt := `INSERT INTO unit_ops (unitop_id, op_key, unit_id,  op_name, op_descr, op_pos) VALUES(?,?,?,?,?,?)`
	_, err := m.DB.ExecContext(ctx, stmt,
		up.ID,
		up.UnitID,
		up.OpKey,
		up.OpName,
		up.OpDescription,
		up.OpPosition,
	)
	if err != nil {
		me, ok := err.(*mysql.MySQLError)
		if !ok {
			return err
		}
		// Eintrag in DB schon vorhanden
		if me.Number == 1062 {
			return nil
		}
	}
	return nil
}

// Speichert einen Wert in die Tabelle
// Parameters:
//   - val: Value Der Eintrag der gespeichert werden soll
//
// Return:
//   - error: Fehlermeldung
func (m *DBModel) SaveUnitValue(val Value) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if val.ID == uuid.Nil {
		val.ID = uuid.New()
	}

	stmt := `INSERT INTO unit_params_values (value_id, unitparam_id, stringvalue, value_set, high, low, cv, unit) VALUES(?,?,?,?,?,?,?,?)`
	_, err := m.DB.ExecContext(ctx, stmt,
		val.ID,
		val.ParamID,
		val.StringValue,
		val.Set,
		val.High,
		val.Low,
		val.Cv,
		val.Unit,
	)
	if err != nil {
		me, ok := err.(*mysql.MySQLError)
		if !ok {
			return err
		}
		// Eintrag in DB schon vorhanden
		if me.Number == 1062 {
			return nil
		}
		return err
	}
	return nil
}
