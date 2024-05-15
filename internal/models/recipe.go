package models

import (
	"context"
	"time"
)

// Speichern eines Rezeptes in der Datenbank
func (m *DBModel) InsertRecipe(recipe Recipe, plantId int) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		INSERT INTO recipes
			(recipe_name, recipe_plantId, created_at, updated_at)
		Values
			(?,?,?,?)
	`
	result, err := m.DB.ExecContext(ctx, stmt,
		recipe.Name,
		plantId,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// kontrolliert ob das Rezept in der DB existiert
func (m *DBModel) ExistRecipe(recipeName string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		SELECT recipe_id
		FROM recipes
		WHERE recipe_name = ?
	`

	row := m.DB.QueryRowContext(ctx, stmt, recipeName)
	err := row.Scan(&recipeName)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Alles Rezepte einer Anlage
func (m *DBModel) GetAllRecipes(plantId int) ([]Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	recipes := []Recipe{}
	stmt := `
	
		SELECT recipe_id, recipe_name, recipe_plantId
		FROM recipes		
		WHERE recipe_plantId = ?
	`

	rows, err := m.DB.QueryContext(ctx, stmt, plantId)
	if err != nil {
		return recipes, err
	}
	defer rows.Close()
	for rows.Next() {
		recipe := Recipe{}
		err = rows.Scan(&recipe.Id, &recipe.Name, &recipe.PlantID)
		if err != nil {
			return recipes, err
		}
		recipes = append(recipes, recipe)
	}
	if err = rows.Err(); err != nil {
		return recipes, err
	}
	return recipes, nil
}
