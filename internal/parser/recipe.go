package parser

type Recipe struct {
	RecipeName string      `json:"recipename"`
	Category   string      `json:"category"`
	Steps      []Step      `json:"steps"`
	Parameters []Parameter `json:"params"`
}
