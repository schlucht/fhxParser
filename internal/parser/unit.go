package parser

type Unit struct {
	UnitName     string      `json:"unitname"`
	UnitCategory string      `json:"unitcategory"`
	UnitPosition string      `json:"unitpositon"`
	Time         int         `json:"time"`
	Author       string      `json:"author"`
	Description  string      `json:"description"`
	Type         string      `json:"type"`
	Parameters   []Parameter `json:"params"`
	Steps        []Step      `json:"steps"`
}
