package fhxModels

type Procedure struct {
	Name       string      `json:"name"`
	Time       int64       `json:"time"`
	Author     string      `json:"author"`
	Parameters []Parameter `json:"params"`
}

