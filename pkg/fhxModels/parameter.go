package fhxModels

type Parameter struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	Value       Value  `json:"value"`
}
