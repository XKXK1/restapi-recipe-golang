package listing

import (
	"errors"
	"math/rand"
)

// ErrNotFound is used when a beer could not be found.
var ErrNotFound = errors.New("recipe not found")

// Repository provides access to the beer and review storage.
type Repository interface {
	// GetBeer returns the beer with given ID.
	GetRecipe(int) (Recipe, error)
	// GetAllBeers returns all beers saved in storage.
	GetAllRecipes() []Recipe
}

// Service provides beer and review listing operations.
type Service interface {
	GetRecipe(int) (Recipe, error)
	GetRecipes() []Recipe
	GetRandomRecipe(string) Recipe
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetBeers returns all beers
func (s *service) GetRecipes() []Recipe {
	return s.r.GetAllRecipes()
}

// GetBeer returns a beer
func (s *service) GetRecipe(id int) (Recipe, error) {
	return s.r.GetRecipe(id)
}

// GetRandomRecipe
func (s *service) GetRandomRecipe(mealType string) Recipe {
	recipes := s.r.GetAllRecipes()
	var outRecipe Recipe

	for outRecipe.ID == 0 {
		ranNumber := rand.Intn(len(recipes))
		if recipes[ranNumber].MealType == mealType {
			outRecipe = recipes[ranNumber]
			break
		}
	}
	return outRecipe
}
