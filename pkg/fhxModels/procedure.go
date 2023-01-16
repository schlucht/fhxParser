package fhxModels

type Procedure struct {
	Name       string      `json:"name"`
	Time       int         `json:"time"`
	Author     string      `json:"author"`
	Parameters []Parameter `json:"params"`
}
