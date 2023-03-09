package fhxModels

type Procedure struct {
	Name        string      `json:"name"`
	Time        int         `json:"time"`
	Author      string      `json:"author"`
	Description string      `json:"desc"`
	Parameters  []Parameter `json:"params"`
}
