package fhx

type Fhx struct {
	FhxId    int               `json:"fhx_id"`
	UnitName string            `json:"unitname"`
	Recipes  []Recipe          `json:"recipes"`
	Units    []Unit            `json:"units"`
	regFhx   map[string]string `json:"-"`
}

type Unit struct {
	UnitId       int         `json:"unit_id"`
	UnitName     string      `json:"unitname"`
	UnitPosition string      `json:"unitpositon"`
	Procedures   []Procedure `json:"unitprocedure"`
}

type Recipe struct {
	RecipeId   int    `json:"recipe_id"`
	RecipeName string `json:"recipename"`
	Steps      []Step `json:"steps"`
}

type Parameter struct {
	ParameterId int    `json:"parameter_id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Value       Value  `json:"value"`
}

type Procedure struct {
	ProcedureId int         `json:"procedure_id"`
	Name        string      `json:"name"`
	Time        int         `json:"time"`
	Author      string      `json:"author"`
	Description string      `json:"desc"`
	Parameters  []Parameter `json:"params"`
}

type Step struct {
	StepId      int         `json:"step_id"`
	Name        string      `json:"name"`
	Key         string      `json:"key"`
	Description string      `json:"desc"`
	Rect        string      `json:"rec"`
	Parameters  []Parameter `json:"parameters"`
}

type Value struct {
	ValueId     int    `json:"value_id"`
	StringValue string `json:"stringvalue,omitempty"`
	Set         string `json:"set,omitempty"`
	High        string `json:"high,omitempty"`
	Low         string `json:"low,omitempty"`
	Cv          string `json:"cv,omitempty"`
	Unit        string `json:"unit,omitempty"`
}
