package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Operation struct {
	ID        uuid.UUID `json:"op_id"`
	OPName    string    `json:"op_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *Operation) IsEmpty() bool {
	return m.ID == uuid.Nil
}

type OperationPlant struct {
	ID          uuid.UUID   `json:"opplant_id"`
	OperationID uuid.UUID   `json:"id_op"`
	PlantID     uuid.UUID   `json:"id_plant"`
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
	OPPlantID   uuid.UUID `json:"id_opplant"`
	Name        string    `json:"param_name"`
	Description string    `json:"param_desc"`
	Value       []Value   `json:"value"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Value struct {
	ID          uuid.UUID `json:"value_id"`
	ParamID     uuid.UUID `json:"id_param"`
	High        int       `json:"high,omitempty"`
	Low         int       `json:"low,omitempty"`
	Cv          int       `json:"cv,omitempty"`
	StringValue string    `json:"stringvalue,omitempty"`
	ValueSet    string    `json:"valueset,omitempty"`
	Unit        string    `json:"unit,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type OPPlant struct {
	OPId          uuid.UUID `json:"op_id"`
	OPPlantID     uuid.UUID `json:"opplant_id"`
	IDPlant       uuid.UUID `json:"id_plant"`
	OpName        string    `json:"op_name"`
	OPCategory    string    `json:"op_category"`
	OPPosition    string    `json:"op_position"`
	OPTime        int       `json:"op_time"`
	OPAuthor      string    `json:"op_author"`
	OPDescription string    `json:"op_descr"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
}

func (m *DBModel) CreateOPTable() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Operation Haupttabelle
	stmt := `
		CREATE TABLE IF NOT EXISTS operations (
			op_id VARCHAR(50) NOT NULL PRIMARY KEY,
			op_name VARCHAR(50) NOT NULL,
			updated_at DATETIME NOT NULL,
			created_at DATETIME NOT NULL,
			CONSTRAINT op_unique UNIQUE (op_name)
		);
	`
	_, err := m.DB.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	// Operation Betrieb Tabelle in Abhängikeit
	// mit der Betriebstabelle und den Operationen
	stmt = `
		CREATE TABLE IF NOT EXISTS op_plant (
			opplant_id VARCHAR(50) NOT NULL PRIMARY KEY,
			id_op VARCHAR(50) NOT NULL,
			id_plant VARCHAR(50) NOT NULL,
			op_category VARCHAR(50) NULL,
			op_position VARCHAR(50) NULL,
			op_author VARCHAR(50) NULL,
			op_desc TEXT NULL,
			op_time INT NULL,
			updated_at DATETIME NOT NULL,
			created_at DATETIME NOT NULL,
			FOREIGN KEY (id_op) REFERENCES operations(op_id),
			FOREIGN KEY (id_plant) REFERENCES plants(plant_id)
		);
	`
	_, err = m.DB.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	// Die Parameter der Operation
	stmt = `
		CREATE TABLE IF NOT EXISTS opparameters (
			params_id VARCHAR(50) NOT NULL PRIMARY KEY,
			id_opplant VARCHAR(50) NOT NULL,
			param_name VARCHAR(50) NOT NULL,
			param_desc TEXT NULL,
			updated_at DATETIME NOT NULL,
			created_at DATETIME NOT NULL,
			FOREIGN KEY (id_opplant) REFERENCES op_plant(opplant_id)
		);
	`
	_, err = m.DB.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	// Parameterwerte speichern. Die Parameter sollten
	// eine Historie bilden.
	stmt = `
	CREATE TABLE IF NOT EXISTS opparamvalues (
		value_id VARCHAR(50) NOT NULL PRIMARY KEY,
		id_params VARCHAR(50) NOT NULL,
		high INT NULL,
		low INT NULL,
		cv INT NULL,
		stringvalue VARCHAR(50) NULL,
		valueset VARCHAR(50) NULL,
		unit VARCHAR(10) NULL,
		param_desc TEXT NULL,
		updated_at DATETIME NOT NULL,
		created_at DATETIME NOT NULL,
		FOREIGN KEY (id_params) REFERENCES opparameters(params_id)
	);
`
	_, err = m.DB.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	return nil
}

// Operation anhand der Betriebs ID auslesen
// Parameter: uuid.UUID
// Return: Operation struct, error
func (m *DBModel) OpFromPlantID(plantId uuid.UUID) ([]OPPlant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		CREATE VIEW IF NOT EXISTS qryOPPlant AS SELECT 	
			op_id,
			opplant_id,
			op.id_plant,
			op_name,
			op.op_category,
			op.op_position,
			op.op_time,
			op.op_author,
			op.op_desc,
			op.updated_at,
			op.created_at 
		FROM operations AS o
		INNER JOIN op_plant AS op ON o.op_id = op.id_op
	`

	_, err := m.DB.ExecContext(ctx, stmt)
	if err != nil {
		return nil, err
	}

	var ops []OPPlant
	stmt = `SELECT 
			*
		FROM 
			qryOPPlant
		WHERE id_plant = ?`
	res, err := m.DB.QueryContext(ctx, stmt, plantId)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var op OPPlant
		err = res.Scan(
			&op.OPId,
			&op.OPPlantID,
			&op.IDPlant,
			&op.OpName,
			&op.OPCategory,
			&op.OPPosition,
			&op.OPTime,
			&op.OPAuthor,
			&op.OPDescription,
			&op.UpdatedAt,
			&op.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		ops = append(ops, op)
	}
	return ops, nil
}

// OP Tabelle einmalig erstellen

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

	stmt := `INSERT INTO operations 
			(op_id, op_name, updated_at, created_at) 
		VALUES
			(?,?,?,?)`
	_, err := m.DB.ExecContext(ctx, stmt,
		op.ID,
		op.OPName,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

// Aktualisiert eine OP in der Datenbank
// Parameter: Operation struct
// Return: error
func (m *DBModel) UpdateOP(op Operation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `UPDATE operations 
		SET op_name = ?, updated_at = ? 
		WHERE op_id = ?`
	_, err := m.DB.ExecContext(ctx, stmt,
		op.OPName,
		time.Now(),
		op.ID,
	)

	if err != nil {
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

	stmt := `SELECT 
			op_id, 
			op_name, 
			created_at, 
			updated_at
		 FROM 
		 	operations 
		WHERE op_id = ?`
	err := m.DB.QueryRowContext(ctx, stmt,
		id.String(),
	).Scan(
		&op.ID,
		&op.OPName,
		&op.CreatedAt,
		&op.UpdatedAt,
	)
	if err != nil {
		return op, err
	}
	return op, nil
}

// Sucht in der Tabelle eine Operation anhand des OPNamen
// Parameter: Name string
// Return: Operation struct, error
func (m *DBModel) OpFromName(name string) (Operation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var op Operation

	stmt := `SELECT 
			op_id, 
			op_name, 
			created_at, 
			updated_at 
		FROM 
			operations 
		WHERE op_name = ?`
	err := m.DB.QueryRowContext(ctx, stmt,
		name,
	).Scan(
		&op.ID,
		&op.OPName,
		&op.CreatedAt,
		&op.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return op, nil
		} else {
			return op, err
		}
	}
	return op, nil
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
		(
			opplant_id, 
			id_op, 
			id_plant, 
			op_category, 
			op_position, 
			op_time, 
			op_author, 
			op_desc,
			updated_at,
			created_at)
		VALUES(?,?,?,?,?,?,?,?,?,?)	`

	_, err := m.DB.ExecContext(ctx, stmt,
		opPlant.ID.String(),
		opPlant.OperationID.String(),
		opPlant.PlantID.String(),
		opPlant.Category,
		opPlant.Position,
		opPlant.OPTime,
		opPlant.Author,
		opPlant.Description,
		time.Now(),
		time.Now(),
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

	stmt := `INSERT INTO opparameters 
			(
				params_id , 
				id_opplant, 
				param_name, 
				param_desc, 
				updated_at, 
				created_at) 
		VALUES
			(?,?,?,?,?,?)`
	_, err := m.DB.ExecContext(ctx, stmt,
		param.ID.String(),
		opPlantID.String(),
		param.Name,
		param.Description,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

// Aktualisieren der Parameter einer OP in der Datenbank
// Parameter:
//   - parser.Parameter param
//   - uuid.UUID opPlantID
//
// Return:
//   - error
func (m *DBModel) UpdateParams(param Parameter, opPlantID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `UPDATE opparameters 
		SET param_name = ?, param_desc = ?, updated_at = ? 
		WHERE id_opplant = ? AND params_id  = ?`
	_, err := m.DB.ExecContext(ctx, stmt,
		param.Name,
		param.Description,
		time.Now(),
		opPlantID.String(),
		param.ID.String(),
	)
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

	stmt := `SELECT 
			opplant_id 
		FROM 
			op_plant 
		WHERE id_op = ? AND id_plant = ?`
	res, err := m.DB.QueryContext(ctx, stmt,
		idOP,
		idPlant,
	)
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

// Gibt die ID einer OP andhand des OP-Namen und der ID der Anlage zurück
//
// Parameters:
//   - name: Name der OP
//   - idPlant: ID der OP-Plant
//
// Return:
//   - uuid.UUID: ID der OP
//   - error: Fehlermeldung
func (m *DBModel) IDOPPlantFromName(name string, idPlant uuid.UUID) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `SELECT 
			opplant_id 
		FROM 
			qryOPPlant 
		WHERE op_name = ? AND id_plant = ?`
	res, err := m.DB.QueryContext(ctx, stmt,
		name,
		idPlant,
	)
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

	stmt := `SELECT 
			params_id 
		FROM 
			opparameters 
		WHERE id_opplant = ? AND param_name = ?`
	res, err := m.DB.QueryContext(ctx, stmt,
		opPlantID.String(),
		paramName,
	)
	if err != nil {
		return uuid.Nil, err
	}
	var uid uuid.UUID
	for res.Next() {
		err = res.Scan(
			&uid,
		)
		if err != nil {
			return uuid.Nil, err
		}
	}
	return uid, nil
}

// Gibt die ID von dem Paramter zurück
//
// Parameter:
//   - string paramname
//   - uuid.UUID ooplantid
//
// Return:
//   - uuid.UUID: ID des Parameters
//   - error: Fehlermeldung
func (m *DBModel) ParamIdFromName(paramname string, ooplantid uuid.UUID) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT 
			params_id 
		FROM 
			opparameters 
		WHERE id_opplant = ? AND param_name = ?`
	res, err := m.DB.QueryContext(ctx, stmt,
		ooplantid.String(),
		paramname,
	)
	if err != nil {
		return uuid.Nil, err
	}
	var uid uuid.UUID
	for res.Next() {
		err = res.Scan(
			&uid,
		)
		if err != nil {
			return uuid.Nil, err
		}
	}
	return uid, nil
}

// Liefert alle Parameter einer OP aus der Datenbank
//
// Parameter:
//   - uuid.UUID ooplantid
//
// Return:
//   - []Parameter: Parameter
//   - error: Fehlermeldung
func (m *DBModel) ParamFromOPPlantID(id uuid.UUID) ([]Parameter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	param := Parameter{}
	params := []Parameter{}

	stmt := `SELECT 
				params_id,
				id_opplant,
				param_name, 
				param_desc,
				updated_at,
				created_at
			FROM opparameters 
			WHERE id_opplant = ? 
			ORDER BY param_name`
	res, err := m.DB.QueryContext(ctx, stmt,
		id.String(),
	)
	for res.Next() {
		res.Scan(
			&param.ID,
			&param.OPPlantID,
			&param.Name,
			&param.Description,
			&param.UpdatedAt,
			&param.CreatedAt,
		)
		if err != nil {
			return params, err
		}
		p, err := m.ValuesFromParamId(param.ID)
		if err != nil {
			return params, err
		}
		param.Value = p
		params = append(params, param)
	}

	return params, nil
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

	stmt := `INSERT INTO 
			opparamvalues 
			(
				value_id, 
				id_params, 
				high, 
				low, 
				cv, 
				unit, 
				stringvalue, 
				valueset,
				updated_at,
				created_at) 
			VALUES
			(?,?,?,?,?,?,?,?,?,?)`
	_, err := m.DB.ExecContext(ctx, stmt,
		value.ID.String(),
		value.ParamID.String(),
		value.High,
		value.Low,
		value.Cv,
		value.Unit,
		value.StringValue,
		value.ValueSet,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

// Neuester Eintrag eines Value von einem Paramater
//
// Parameter:
//   - uuid.UUID paramId
//
// Return:
//   - Value: Value
//   - error: Fehlermeldung
func (m *DBModel) LastValueFromParamId(paramId uuid.UUID) (Value, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	val := Value{}
	stmt := `SELECT 
			value_id, 
			unitparams_id, 
			high, 
			low, 
			cv, 
			unit, 
			stringvalue, 
			valueset,
			updated_at,
			created_at
		FROM 
			unitparameters_values 
		WHERE unitparams_id = ? 
		ORDER BY value_id DESC
		LIMIT 1`

	err := m.DB.QueryRowContext(ctx, stmt,
		paramId.String(),
	).Scan(
		&val.ID,
		&val.ParamID,
		&val.High,
		&val.Low,
		&val.Cv,
		&val.Unit,
		&val.StringValue,
		&val.ValueSet,
		&val.UpdatedAt,
		&val.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return val, nil
		}
		return val, err
	}
	return val, nil
}

func (m *DBModel) ValuesFromParamId(id uuid.UUID) ([]Value, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	val := Value{}
	values := []Value{}

	stmt := `SELECT 
			value_id, 
			id_params, 
			high, 
			low, 
			cv, 
			unit, 
			stringvalue, 
			valueset,
			updated_at,
			created_at
		FROM 
			opparamvalues 
		WHERE params_id = ? 
		ORDER BY value_id DESC`

	res, err := m.DB.QueryContext(ctx, stmt, id.String())

	for res.Next() {
		res.Scan(
			&val.ID,
			&val.ParamID,
			&val.High,
			&val.Low,
			&val.Cv,
			&val.Unit,
			&val.StringValue,
			&val.ValueSet,
			&val.UpdatedAt,
			&val.CreatedAt,
		)
		if err != nil {
			return values, err
		}
		values = append(values, val)
	}
	return values, nil
}
