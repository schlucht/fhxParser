package fhxModels

type Step struct {
	Name       string    `json:"name"`
	Key        string    `json:"key"`
	Author     string    `json:"author"`
	Date       int       `json:"date"`
	Parameters Parameter `json:"parameters"`
}
