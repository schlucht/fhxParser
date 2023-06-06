package models

type Recipe struct {	
	RecipeName string `json:"recipename"`
	Steps      []Step `json:"steps"`
}
