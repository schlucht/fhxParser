package fhxModels

type Step struct {
	Name        string      `json:"name"`
	Key         string      `json:"key"`
	Description string      `json:"desc"`
	Rect        string      `json:"rec"`
	Parameters  []Parameter `json:"parameters"`
}
