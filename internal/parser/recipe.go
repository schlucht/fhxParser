package parser

type Recipe struct {
	RecipeName string      `json:"recipename"`
	Steps      []Step      `json:"steps"`
	Parameters []Parameter `json:"params"`
}
