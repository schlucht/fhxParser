package fhx

import "time"

type Fhx struct {
	FhxId   int               `json:"fhx_id"`
	Recipes []Recipe          `json:"recipes,omitempty"`
	Units   []Unit            `json:"units,omitempty"`
	regFhx  map[string]string `json:"-"`
}

type Unit struct {
	UnitId       int         `json:"unit_id"`
	UnitName     string      `json:"unitname"`
	UnitPosition string      `json:"unitpositon"`
	Time         int         `json:"time"`
	Author       string      `json:"author"`
	Description  string      `json:"description"`
	Parameters   []Parameter `json:"params"`
}

type Parameter struct {
	ParameterId int     `json:"parameter_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Value       []Value `json:"value"`
}

type Value struct {
	ValueId     int       `json:"value_id"`
	StringValue string    `json:"stringvalue,omitempty"`
	Set         string    `json:"value_set,omitempty"`
	High        int       `json:"high,omitempty"`
	Low         int       `json:"low,omitempty"`
	Cv          int       `json:"cv,omitempty"`
	Unit        string    `json:"unit,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type Recipe struct {
	RecipeId   int    `json:"recipe_id"`
	RecipeName string `json:"recipename"`
	Steps      []Step `json:"steps"`
}

type Step struct {
	StepId      int         `json:"step_id"`
	Name        string      `json:"name"`
	Key         string      `json:"key"`
	Description string      `json:"description"`
	Rect        string      `json:"rec"`
	Parameters  []Parameter `json:"parameters"`
}
