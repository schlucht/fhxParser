package parser

type Operation struct {
	UnitName     string      `json:"Opname"`
	UnitPosition string      `json:"Oppositon"`
	Time         int         `json:"time"`
	Author       string      `json:"author"`
	Description  string      `json:"description"`
	Type         string      `json:"type"`
	Parameters   []Parameter `json:"params"`
}
