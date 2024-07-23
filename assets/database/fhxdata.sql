

-- Create Operationstable
CREATE TABLE operations (
    op_id VARCHAR(50) PRIMARY KEY,
    opname VARCHAR(50),
    updated_at TIMESTAMP,
    created_at TIMESTAMP
);

-- Create Operation TO Plant Table
CREATE TABLE op_plant (
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
);

-- Create Unit Table to Plant
CREATE TABLE units (
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
);

-- Create UNIT_OPS to Operations
CREATE TABLE unit_ops (
    unitop_id VARCHAR(50)  PRIMARY KEY,
    unit_id VARCHAR(50),
    op_key VARCHAR(50),
    op_descr TEXT,
    op_pos VARCHAR(50),
    updated_at TIMESTAMP,
    created_at TIMESTAMP,
    FOREIGN KEY (unitop_id) REFERENCES units(unit_id)
);

-- Create Unitparameters
CREATE TABLE unitparameters (
    unitparam_id VARCHAR(50)  PRIMARY KEY,
    unitop_id VARCHAR(50),
    originValue_id VARCHAR(50),
    param_name VARCHAR(255),
    origin VARCHAR(50),
    deferto VARCHAR(50),
    updated_at TIMESTAMP,
    created_at TIMESTAMP,
    FOREIGN KEY (unitop_id) REFERENCES unit_op(unitop_id)
);

CREATE TABLE unitparameters_values (
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
)

-- CREATE recipes
CREATE TABLE recipes (
    recipe_id VARCHAR(50)  PRIMARY KEY,
    plant_id VARCHAR(50),
    recipe_name VARCHAR(100),
    recipe_category VARCHAR(50),
    updated_at TIMESTAMP,
    created_at TIMESTAMP,
    FOREIGN KEY (plant_id) REFERENCES plants(plant_id)
);

--  CREATE recipe_steps
CREATE TABLE recipe_steps (
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
)

-- CREATE recipe stepattributes
CREATE TABLE recipe_stepattributes (
    stepparam_id VARCHAR(50)  PRIMARY KEY,
    unitparam_id VARCHAR(50),
    descr TEXT,
    updated_at TIMESTAMP,
    created_at TIMESTAMP,
    FOREIGN KEY (unitparam_id) REFERENCES unitparameters(unitparam_id)
)

-- Create Rezept Values
CREATE TABLE recipe_values (
    value_id VARCHAR(50) PRIMARY KEY,
    stepattribute_id VARCHAR(50),
    cv INTEGER,
    stringvalue VARCHAR(50),
    value_set VARCHAR(50),
    updated_at TIMESTAMP,
    created_at TIMESTAMP,
     FOREIGN KEY (stepattribute_id) REFERENCES recipe_stepattributes(stepparam_id)
)