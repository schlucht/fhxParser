package models

import (
	"context"
	"fmt"
	"time"

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
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UnitParameter struct {
	ID            uuid.UUID `json:"unitparam_id"`
	UnitOPID      uuid.UUID `json:"unitop_id"`
	OriginValueID uuid.UUID `json:"originValue_id"`
	ParamName     string    `json:"param_name"`
	Deferto       string    `json:"deferto"`
	Origin        string    `json:"origin"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
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
func (m *DBModel) SaveUnit(up Unit) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if up.ID == uuid.Nil {
		up.ID = uuid.New()
	}

	stmt := `INSERT INTO units 
			(unit_id, plant_id, unit_name, unit_category, unit_pos, unit_time, unit_author, unit_descr, created_at, updated_at) 
		VALUES
			(?,?,?,?,?,?,?,?,?,?)`
	_, err := m.DB.ExecContext(ctx, stmt,
		up.ID,
		up.PlantID,
		up.UnitName,
		up.UnitCategory,
		up.UnitPosition,
		up.Time,
		up.Author,
		up.Description,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

// Aktualisiert eine Unit in der Datenbank
// Parameter: Unit struct
// Return: error
func (m *DBModel) UpdateUnit(up Unit) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `UPDATE units 
		SET 
			unit_name = ?, 
			unit_category = ?, 
			unit_pos = ?, 
			unit_time = ?, 
			unit_author = ?, 
			unit_descr = ?, 
			updated_at = ?
		WHERE unit_id = ?`
	_, err := m.DB.ExecContext(ctx, stmt,
		up.UnitName,
		up.UnitCategory,
		up.UnitPosition,
		up.Time,
		up.Author,
		up.Description,
		time.Now(),
		up.ID,
	)
	if err != nil {
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

	stmt := `SELECT 
			unit_id 
		FROM units 
		WHERE plant_id = ? AND unit_name = ?`
	var id uuid.UUID
	row := m.DB.QueryRowContext(ctx, stmt, plantID, upName)
	err := row.Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

// Alle Unit einer Anlage auslesen
//
// Parameters:
//   - plantID: ID der Anlage
//
// Return:
//   - []Unit: Liste der Units
//   - error: Fehlermeldung
func (m *DBModel) UnitIdFromPlantId(plantID uuid.UUID) ([]Unit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var unit Unit
	var units = []Unit{}

	stmt := `SELECT 
			unit_id,
			plant_id,
			unit_name,
			unit_category,
			unit_pos,
			unit_time,
			unit_author,
			unit_descr,
			created_at,
			updated_at
		FROM units 
		WHERE plant_id = ?`

	row, err := m.DB.QueryContext(ctx, stmt, plantID.String())
	if err != nil {
		return units, err
	}
	for row.Next() {
		err = row.Scan(
			&unit.ID,
			&unit.PlantID,
			&unit.UnitName,
			&unit.UnitCategory,
			&unit.UnitPosition,
			&unit.Time,
			&unit.Author,
			&unit.Description,
			&unit.CreatedAt,
			&unit.UpdatedAt,
		)
		if err != nil {
			return units, err
		}
		units = append(units, unit)
	}	
	return units, nil
}

// Gibt die ID einer OP mit der Unit ID zurück
// Anhand des OP-Namen
//
// Parameters:
//   - opName: Name der OP
//   - unitID: ID der Unit
//
// Return:
//   - uuid.UUID: ID der OP
//   - error: Fehlermeldung
func (m *DBModel) OpUnitIdFromName(opName string, unitID uuid.UUID) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT 
			unitop_id 
		FROM unit_ops 
		WHERE unit_id = ? AND op_name = ?`

	var id uuid.UUID
	row := m.DB.QueryRowContext(ctx, stmt, unitID, opName)
	err := row.Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

// Gibt die ID einer OP mit der Unit ID zurück
// Anhand des OP-Keys
//
// Parameters:
//   - opKey: Name der OP
//   - unitID: ID der Unit
//
// Return:
//   - uuid.UUID: ID der OP
//   - error: Fehlermeldung
func (m *DBModel) OpUnitIdFromKey(opKey string, unitID uuid.UUID) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT 
			unit_id 
		FROM unit_ops 
		WHERE unit_id = ? AND op_key = ?`

	var id uuid.UUID
	row := m.DB.QueryRowContext(ctx, stmt, unitID, opKey)
	err := row.Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

// Speichert einen Wert in die Tabelle
// Parameters: UnitOP struct
// Return: error
func (m *DBModel) SaveUnitOps(up UnitOP) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if up.ID == uuid.Nil {
		up.ID = uuid.New()
	}

	stmt := `INSERT INTO unit_ops 
			(unitop_id, unit_id, op_key_id,  op_name, op_descr, op_pos, created_at, updated_at) 
		VALUES
			(?,?,?,?,?,?,?,?)`
	_, err := m.DB.ExecContext(ctx, stmt,
		up.ID,
		up.UnitID,
		up.OpKey,
		up.OpName,
		up.OpDescription,
		up.OpPosition,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

// Aktualisiert eine OP in der Datenbank
// Parameter: UnitOP struct
// Return: error
func (m *DBModel) UpdateUnitOP(up UnitOP) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `UPDATE unit_ops 
		SET 
			op_name = ?, 
			op_descr = ?, 
			op_pos = ?, 
			updated_at = ?
		WHERE unitop_id = ?`
	_, err := m.DB.ExecContext(ctx, stmt,
		up.OpName,
		up.OpDescription,
		up.OpPosition,
		time.Now(),
		up.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

// Speichert einen Wert in die Tabelle
// Parameters:
//   - val: Value Der Eintrag der gespeichert werden soll
//
// Return:
//   - error: Fehlermeldung
func (m *DBModel) SaveUnitParameter(param UnitParameter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if param.ID == uuid.Nil {
		param.ID = uuid.New()
	}
	stmt := `INSERT INTO unitparameters
			(
				unitparam_id, 
				unitop_id, 
				originValue_id, 
				param_name, 
				origin, 
				deferto,
				created_at,
				updated_at
				) 
		VALUES
			(?,?,?,?,?,?,?,?)`
	_, err := m.DB.ExecContext(ctx, stmt,
		param.ID,
		param.UnitOPID,
		param.OriginValueID,
		param.ParamName,
		param.Origin,
		param.Deferto,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

// Prüft ob der UnitParameter schon in der Datenbank existiert
// Parameter: UnitParameter struct
// Return: uuid.UUID, error
func (m *DBModel) ExistUnitParam(unitopId uuid.UUID, paramName string) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `SELECT 
			unitparam_id 
		FROM unitparameters 
		WHERE unitop_id = ? AND param_name = ?`
	var id uuid.UUID
	row := m.DB.QueryRowContext(ctx, stmt, unitopId, paramName)
	err := row.Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

// Gibt die ID von dem Paramter zurück nach der Deferto ID
// Parameter:
//   - UnitParameter struct
//   - deferName: Name des
//
// Return:
//   - uuid.UUID,
//   - error Fehlermeldung
func (m *DBModel) IDUnitParamDeferTo(unitopId uuid.UUID, deferName string) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT 
			unitparam_id 
		FROM unitparameters 
		WHERE unitop_id = ? AND deferto = ?`
	var id uuid.UUID
	row := m.DB.QueryRowContext(ctx, stmt, unitopId, deferName)
	err := row.Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

// Speichert die Values eines Parameters in die Datenbank
// Die Values sind nur die Werte für die Deferred Paramater
// Die Orginal Values sind in den Unitparameters gespeichert
// Parameter:
//   - val: UnitValue struct
//
// Return:
//   - error: Fehlermeldung
func (m *DBModel) SaveUnitParamValue(val UnitValue) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if val.ID == uuid.Nil {
		val.ID = uuid.New()
	}
	fmt.Println(val.ID, val.UnitID)
	stmt := `INSERT INTO unitparameters_values
			(value_id, 
				unitparams_id,
				high,
				low,
				cv,
				unit,
				stringvalue,
				valueset,
				created_at,
				updated_at
				) 
		VALUES
			(?,?,?,?,?,?,?,?,?,?)`
	_, err := m.DB.ExecContext(ctx, stmt,
		uuid.New(),
		val.UnitID,
		val.High,
		val.Low,
		val.Cv,
		val.Unit,
		val.StringValue,
		val.Set,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}
