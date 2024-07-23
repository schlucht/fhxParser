package models

type System struct {
}

func (m *DBModel) GetSystemTable() {
	tbls, err := m.DB.Query("SELECT COUNT(*) FROM duckdb_tables();")
	if err != nil {
		m.errorLog.Println(err)
	}
	var cols int

	for tbls.Next() {
		if err = tbls.Scan(&cols); err != nil {
			m.errorLog.Println(err)
		}
	}
	m.infoLog.Println(cols)
}

func (m *DBModel) createPlantTable() error {
	sql := `CREATE TABLE IF NOT EXISTS plants (
		plant_id VARCHAR(50) PRIMARY KEY,
		plant VARCHAR(20),
		updated_at TIMESTAMP,
		created_at TIMESTAMP
		);`
	err := m.CreateTable(sql)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) createOPPlantTable() {
	sql := `CREATE TABLE op_plant (
		opplant_id VARCHAR(50) PRIMARY KEY,
		id_op VARCHAR(50),
		id_plant VARCHAR(50),
		op_category VARCHAR(50),
		op_position VARCHAR(50),
		op_time INTEGER,
		op_author VARCHAR(50),
		op_description TEXT,
		updated_at TIMESTAMP,
		created_at TIMESTAMP,
		FOREIGN KEY (id_op) REFERENCES operations(op_id),
		FOREIGN KEY (id_plant) REFERENCES plants(plant_id)
	);`
	err := m.CreateTable(sql)
	if err != nil {
		m.errorLog.Panic(err)
	}
}

func (m *DBModel) createUnitsTable() {
	sql := `CREATE TABLE units (
		unit_id VARCHAR(50) PRIMARY KEY,
		plant_id VARCHAR(50),
		unit_name VARCHAR(255),
		unit_category VARCHAR(50),
		unit_pos VARCHAR(50),
		unit_time INTEGER,
		unit_author VARCHAR(50),
		unit_descr TEXT,
		updated_at TIMESTAMP,
		created_at TIMESTAMP,
		FOREIGN KEY (plant_id) REFERENCES plants(plant_id)
	);`
	err := m.CreateTable(sql)
	if err != nil {
		m.errorLog.Panic(err)
	}
}

func (m *DBModel) createUnitOPTable() {
	sql := `CREATE TABLE unit_ops (
		unitop_id VARCHAR(50)  PRIMARY KEY,
		unit_id VARCHAR(50),
		op_key VARCHAR(50),
		op_descr TEXT,
		op_pos VARCHAR(50),
		updated_at TIMESTAMP,
		created_at TIMESTAMP,
		FOREIGN KEY (unitop_id) REFERENCES units(unit_id)
	);`
	err := m.CreateTable(sql)
	if err != nil {
		m.errorLog.Panic(err)
	}

}

func (m *DBModel) createUnitParamatersTable() {
	sql := `CREATE TABLE unitparameters (
		unitparam_id VARCHAR(50)  PRIMARY KEY,
		unitop_id VARCHAR(50),
		originValue_id VARCHAR(50),
		param_name VARCHAR(255),
		origin VARCHAR(50),
		deferto VARCHAR(50),
		updated_at TIMESTAMP,
		created_at TIMESTAMP,
		FOREIGN KEY (unitop_id) REFERENCES unit_op(unitop_id)
	);`
	err := m.CreateTable(sql)
	if err != nil {
		m.errorLog.Panic(err)
	}
}

func (m *DBModel) createUnitParameterValueTable() {
	sql := `CREATE TABLE unitparameters_values (
		value_id VARCHAR(50)  PRIMARY KEY,
		unitparams_id VARCHAR(50),
		high INTEGER,
		low INTEGER,
		cv INTEGER,
		unit VARCHAR(10),
		stringvalue VARCHAR(50),
		valueset VARCHAR(50),
		updated_at TIMESTAMP,
		created_at TIMESTAMP,
		FOREIGN KEY (unitparams_id) REFERENCES unitparamaters(unitparam_id)
	)`
	err := m.CreateTable(sql)
	if err != nil {
		m.errorLog.Panic(err)
	}
}

func (m *DBModel) createRecipeTable() {
	sql := `CREATE TABLE recipes (
		recipe_id VARCHAR(50)  PRIMARY KEY,
		plant_id VARCHAR(50),
		recipe_name VARCHAR(100),
		recipe_category VARCHAR(50),
		updated_at TIMESTAMP,
		created_at TIMESTAMP,
		FOREIGN KEY (plant_id) REFERENCES plants(plant_id)
	);`
	err := m.CreateTable(sql)
	if err != nil {
		m.errorLog.Panic(err)
	}
}

func (m *DBModel) createRecipeStepsTable() {
	sql := `CREATE TABLE recipe_steps (
		step_id VARCHAR(50)  PRIMARY KEY,
		recipe_id VARCHAR(50),
		up_id VARCHAR(50),
		up_key VARCHAR(50),
		recipe_descr TEXT,
		recipe_pos VARCHAR(50),
		updated_at TIMESTAMP,
		created_at TIMESTAMP,
		FOREIGN KEY (recipe_id) REFERENCES recipes(recipe_id),
		FOREIGN KEY (up_id) REFERENCES units(unit_id)
	)`
	err := m.CreateTable(sql)
	if err != nil {
		m.errorLog.Panic(err)
	}
}

func (m *DBModel) createRecipeAttributeTable() {
	sql := `CREATE TABLE recipe_stepattributes (
		stepparam_id VARCHAR(50)  PRIMARY KEY,
		unitparam_id VARCHAR(50),
		descr TEXT,
		updated_at TIMESTAMP,
		created_at TIMESTAMP,
		FOREIGN KEY (unitparam_id) REFERENCES unitparameters(unitparam_id)
	)`
	err := m.CreateTable(sql)
	if err != nil {
		m.errorLog.Panic(err)
	}
}

func (m *DBModel) createRecipeValueTable() {
	sql := `CREATE TABLE recipe_values (
		value_id VARCHAR(50) PRIMARY KEY,
		stepattribute_id VARCHAR(50),
		cv INTEGER,
		stringvalue VARCHAR(50),
		value_set VARCHAR(50),
		updated_at TIMESTAMP,
		created_at TIMESTAMP,
		FOREIGN KEY (stepattribute_id) REFERENCES recipe_stepattributes(stepparam_id)
	)`
	err := m.CreateTable(sql)
	if err != nil {
		m.errorLog.Panic(err)
	}
}
