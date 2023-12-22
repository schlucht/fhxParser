package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

// for Database connection
type DBModel struct {
	DB       *sql.DB
	infoLog  *log.Logger
	errorLog *log.Logger
}

// Wrapper for all Models
type Models struct {
	DB DBModel
}

// New Models return a Model
func NewModel(db *sql.DB) Models {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return Models{
		DB: DBModel{
			DB:       db,
			infoLog:  infoLog,
			errorLog: errorLog,
		},
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

type Recipe struct {
	Id        int       `json:"id"`
	Name      string    `json:"recipe_name"`
	Author    string    `json:"desc"`
	UnitID    int       `json:"unit_id"`
	PlantID   int       `json:"plant_id"`
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
	Value       Value     `json:"val"`
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
