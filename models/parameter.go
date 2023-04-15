package models

type Parameter struct {
	ParameterId int     `json:"parameter_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Value       []Value `json:"value"`
}
