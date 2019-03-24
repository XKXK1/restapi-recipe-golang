package memory

import (
	"../../adding"
	"../../listing"
)

type Storage struct {
	recipes []Recipe
}

func (m *Storage) AddRecipe(r adding.Recipe) error {

	newR := Recipe{
		ID:        len(m.recipes) + 1,
		MealType:  r.MealType,
		Name:      r.Name,
		Ingredients: r.Ingredients,
		Preparation: r.Preparation,
	}
	m.recipes = append(m.recipes, newR)

	return nil
}

// Get returns a recipe with the specified ID
func (m *Storage) GetRecipe(id int) (listing.Recipe) {
	var recipe listing.Recipe

	for i := range m.recipes {

		if m.recipes[i].ID == id {
			recipe.ID = m.recipes[i].ID
			recipe.MealType = m.recipes[i].MealType
			recipe.Name = m.recipes[i].Name
			recipe.Ingredients = m.recipes[i].Ingredients
			recipe.Preparation = m.recipes[i].Preparation

			return recipe
		}
	}

	return recipe
}

// GetAll return all recipes
func (m *Storage) GetAllRecipes() []listing.Recipe {
	var recipes []listing.Recipe

	for i := range m.recipes {

		recipe := listing.Recipe{
			ID:        m.recipes[i].ID,
			MealType:  m.recipes[i].MealType,
			Name:      m.recipes[i].Name,
			Ingredients: m.recipes[i].Ingredients,
			Preparation: m.recipes[i].Preparation,
			
		}

		recipes = append(recipes, recipe)
	}

	return recipes
}