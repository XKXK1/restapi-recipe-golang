package adding



// Service provides recipe adding operations.
type Service interface {
	AddRecipe(...Recipe)
}

// Repository provides access to recipe repository.
type Repository interface {
	// AddRecipe saves a given recipe to the repository.
	AddRecipe(Recipe) error
}

type service struct {
	rR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddRecipe adds the given recipe(s) to the database
func (s *service) AddRecipe(r ...Recipe) {

	// any validation can be done here

	for _, recipe := range r {
		_ = s.rR.AddRecipe(recipe) // error handling omitted for simplicity
	}
}

