package models

import (
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

type Plant struct {
	Id        int       `json:"id"`
	Plant     string    `json:"plant_name"`
	CreatetAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Typ struct {
	Id        int       `json:"id"`
	TypeName  string    `json:"type_name"`
	CreatetAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Unit struct {
	Id          int       `json:"id"`
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	Position    string    `json:"position"`
	Author      string    `json:"author"`
	Description string    `json:"desc"`
	Time        int       `json:"time"`
	PlantID     int       `json:"plant_id"`
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

type Operation struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Plant       string `json:"plant"`
	Author      string `json:"author"`
	Description string `json:"desc"`
	Plant_id    int    `json:"-"`
}
