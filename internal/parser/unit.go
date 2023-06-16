package parser

type Unit struct {	
	UnitName     string      `json:"unitname"`
	UnitPosition string      `json:"unitpositon"`
	Time         int         `json:"time"`
	Author       string      `json:"author"`
	Description  string      `json:"description"`
	Parameters   []Parameter `json:"params"`
}
