package models

type Unit struct {
	UnitId       int         `json:"unit_id"`
	UnitName     string      `json:"unitname"`
	UnitPosition string      `json:"unitpositon"`
	Time         int         `json:"time"`
	Author       string      `json:"author"`
	Description  string      `json:"description"`
	Parameters   []Parameter `json:"params"`
}
