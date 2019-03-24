package memory

import (
	"../../adding"
	"../../listing"
)

// Storage stores all recipes in a slice
type Storage struct {
	recipes []Recipe
}

// AddRecipe adds a recipe
func (m *Storage) AddRecipe(r adding.Recipe) error {

	newR := Recipe{
		ID:          len(m.recipes) + 1,
		MealType:    r.MealType,
		Name:        r.Name,
		Ingredients: r.Ingredients,
		Preparation: r.Preparation,
	}
	m.recipes = append(m.recipes, newR)

	return nil
}

// GetRecipe returns a recipe with the specified ID
func (m *Storage) GetRecipe(id int) (listing.Recipe, error) {
	var recipe listing.Recipe

	for i := range m.recipes {

		if m.recipes[i].ID == id {
			recipe.ID = m.recipes[i].ID
			recipe.MealType = m.recipes[i].MealType
			recipe.Name = m.recipes[i].Name
			recipe.Ingredients = m.recipes[i].Ingredients
			recipe.Preparation = m.recipes[i].Preparation

			return recipe, nil
		}
	}

	return recipe, listing.ErrNotFound
}

// GetAllRecipes return all recipes
func (m *Storage) GetAllRecipes() []listing.Recipe {
	var recipes []listing.Recipe

	for i := range m.recipes {

		recipe := listing.Recipe{
			ID:          m.recipes[i].ID,
			MealType:    m.recipes[i].MealType,
			Name:        m.recipes[i].Name,
			Ingredients: m.recipes[i].Ingredients,
			Preparation: m.recipes[i].Preparation,
		}

		recipes = append(recipes, recipe)
	}

	return recipes
}

// AddSampleRecipes adds Sample Recipes
func (m *Storage) AddSampleRecipes() {

	for _, recipe := range SampleMeals {
		m.recipes = append(m.recipes, recipe)
	}

}

// DeleteRecipe deletes a recipe by given id
func (m *Storage) DeleteRecipe(ID int) {
	for i, recipe := range m.recipes {
		if ID == recipe.ID {
			m.recipes = append(m.recipes[:i], m.recipes[i+1:]...)
			break
		}
	}
}
