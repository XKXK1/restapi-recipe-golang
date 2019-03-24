package listing


// Repository provides access to the beer and review storage.
type Repository interface {
	// GetBeer returns the beer with given ID.
	GetRecipe(int) (Recipe)
	// GetAllBeers returns all beers saved in storage.
	GetAllRecipes() []Recipe
}

// Service provides beer and review listing operations.
type Service interface {
	GetRecipe(int) (Recipe)
	GetRecipes() []Recipe
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
func (s *service) GetRecipe(id int) (Recipe) {
	return s.r.GetRecipe(id)
}
