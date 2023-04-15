package models

type Recipe struct {
	RecipeId   int    `json:"recipe_id"`
	RecipeName string `json:"recipename"`
	Steps      []Step `json:"steps"`
}
