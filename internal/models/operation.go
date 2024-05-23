package models

import (
	"context"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Operation struct {
	ID        uuid.UUID `json:"op_id"`
	OPName    string    `json:"opname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OperationPlant struct {
	ID          uuid.UUID   `json:"opplant_id"`
	PlantID     uuid.UUID   `json:"plant_id"`
	OperationID uuid.UUID   `json:"operation_id"`
	Category    string      `json:"op_category"`
	Position    string      `json:"op_position"`
	Author      string      `json:"op_author"`
	Description string      `json:"op_description"`
	OPTime      int         `json:"op_time"`
	Parameters  []Parameter `json:"params"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type Parameter struct {
	ID          uuid.UUID `json:"params_id"`
	OPPlantID   uuid.UUID `json:"opplant_id"`
	Name        string    `json:"params_name"`
	Description string    `json:"params_descr"`
	Value       []Value   `json:"value"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Value struct {
	ID          uuid.UUID `json:"value_id"`
	ParamID     uuid.UUID `json:"params_id"`
	StringValue string    `json:"stringvalue,omitempty"`
	Set         string    `json:"value_set,omitempty"`
	High        int       `json:"high,omitempty"`
	Low         int       `json:"low,omitempty"`
	Cv          int       `json:"cv,omitempty"`
	Unit        string    `json:"unit,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateNewOP inserts a new operation into the operations table in the database.
//
// It takes an Operation struct as a parameter, which contains the op_id, opname, created_at, and updated_at fields.
// The function returns an error if there was a problem executing the SQL statement.
func (m *DBModel) NewOP(op Operation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if op.ID == uuid.Nil {
		op.ID = uuid.New()
	}

	stmt := `INSERT INTO operations (op_id, opname) VALUES(?,?)`
	_, err := m.DB.ExecContext(ctx, stmt, op.ID, op.OPName)
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

// OPFromID retrieves an operation from the database based on its ID.
//
// It takes a UUID as a parameter representing the ID of the operation to retrieve.
// The function returns an Operation struct and an error. The Operation struct contains the op_id, opname, created_at, and updated_at fields of the operation.
// If the operation is found in the database, it is returned along with a nil error.
// If the operation is not found, an empty Operation struct is returned along with an error indicating that the operation was not found.
func (m *DBModel) OPFromID(id uuid.UUID) (Operation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var op Operation

	stmt := `SELECT op_id, opname, created_at, updated_at FROM operations WHERE op_id = ?`
	err := m.DB.QueryRowContext(ctx, stmt, id.String()).Scan(&op.ID, &op.OPName, &op.CreatedAt, &op.UpdatedAt)
	if err != nil {
		return op, err
	}
	return op, nil
}

// OpFromName retrieves the UUID of an operation from the database based on its name.
//
// It takes a string parameter representing the name of the operation to retrieve.
// The function returns a UUID and an error. The UUID represents the ID of the operation.
// If the operation is found in the database, the UUID is returned along with a nil error.
// If the operation is not found, a UUID with a value of uuid.Nil is returned along with an error indicating that the operation was not found.
func (m *DBModel) OpFromName(name string) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT op_id FROM operations WHERE opname = ?`
	res, err := m.DB.QueryContext(ctx, stmt, name)
	if err != nil {
		return uuid.Nil, err
	}

	var uid uuid.UUID
	for res.Next() {
		err = res.Scan(&uid)
		if err != nil {
			return uuid.Nil, err
		}
	}
	return uid, nil
}

// NewOPPlant inserts a new operation plant into the op_plant table in the database.
//
// It takes an OperationPlant struct as a parameter, which contains the opplant_id, id_op, id_plant, op_category, op_position, op_time, op_author, and op_description fields.
// The function returns an error if there was a problem executing the
func (m *DBModel) NewOPPlant(opPlant OperationPlant) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if opPlant.ID == uuid.Nil {
		opPlant.ID = uuid.New()
	}

	stmt := `
		INSERT INTO op_plant
		(opplant_id, id_op, id_plant, op_category, op_position, op_time, op_author, op_description)
		VALUES(?,?,?,?,?,?,?,?)	`

	_, err := m.DB.ExecContext(ctx, stmt,
		opPlant.ID.String(),
		opPlant.OperationID.String(),
		opPlant.PlantID.String(),
		opPlant.Category,
		opPlant.Position,
		opPlant.OPTime,
		opPlant.Author,
		opPlant.Description,
	)
	if err != nil {
		return err
	}
	return nil
}

// NewParam inserts a new parameter into the parameters table in the database.
//
// It takes a Parameter struct and a UUID representing the ID of the operation plant as parameters.
// The function returns an error if there was a problem executing the SQL statement.
func (m *DBModel) NewParam(param Parameter, opPlantID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if param.ID == uuid.Nil {
		param.ID = uuid.New()
	}

	stmt := `INSERT INTO parameters (params_id, opplant_id, params_name, params_descr) VALUES(?,?,?,?)`
	_, err := m.DB.ExecContext(ctx, stmt, param.ID.String(), opPlantID.String(), param.Name, param.Description)
	if err != nil {
		return err
	}
	return nil
}

// ParamFromID retrieves the opplant_id_id from the op_plant table in the database based on the provided idOP and idPlant.
//
// Parameters:
// - idOP: the UUID of the id_op to search for.
// - idPlant: the UUID of the id_plant to search for.
//
// Returns:
// - uuid.UUID: the opplant_id_id found in the op_plant table.
// - error: an error if the query execution or scanning fails.
func (m *DBModel) OPPlantFromID(idOP uuid.UUID, idPlant uuid.UUID) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT opplant_id FROM op_plant WHERE id_op = ? AND id_plant = ?`
	res, err := m.DB.QueryContext(ctx, stmt, idOP, idPlant)
	if err != nil {
		return uuid.Nil, err
	}
	var uid uuid.UUID
	for res.Next() {
		err = res.Scan(&uid)
		if err != nil {
			return uuid.Nil, err
		}
	}
	return uid, nil
}

// ExistParam checks if a parameter with the given name exists for a specific operation plant in the database.
//
// Parameters:
// - opPlantID: the UUID of the operation plant to search for.
// - paramName: the name of the parameter to search for.
//
// Returns:
// - error: an error if there was a problem executing the SQL statement or scanning the result.
func (m *DBModel) ExistParam(opPlantID uuid.UUID, paramName string) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT params_id FROM parameters WHERE opplant_id = ? AND params_name = ?`
	res, err := m.DB.QueryContext(ctx, stmt, opPlantID.String(), paramName)
	if err != nil {
		return uuid.Nil, err
	}
	var uid uuid.UUID
	for res.Next() {
		err = res.Scan(&uid)
		if err != nil {
			return uuid.Nil, err
		}
	}
	return uid, nil
}

// NewValue inserts a new value into the values table in the database.
//
// It takes a Value struct and a UUID representing the ID of the parameter as parameters.
// The function returns an error if there was a problem executing the SQL statement.
//
// Parameters:
// - value: The Value struct containing the values to be inserted.
// - paramID: The UUID representing the ID of the parameter.
//
// Return:
// - error: An error if there was a problem executing the SQL statement, otherwise nil.
func (m *DBModel) NewValue(value Value) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if value.ID == uuid.Nil {
		value.ID = uuid.New()
	}

	stmt := `INSERT INTO paramvalues (value_id, params_id, high, low, cv, unit, stringvalue, valueset) VALUES(?,?,?,?,?,?,?,?)`
	_, err := m.DB.ExecContext(ctx, stmt,
		value.ID.String(),
		value.ParamID.String(),
		value.High,
		value.Low,
		value.Cv,
		value.Unit,
		value.StringValue,
		value.Set)
	if err != nil {
		return err
	}
	return nil
}
