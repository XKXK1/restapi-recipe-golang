package deleting

// Service interface defines methods which can be used
type Service interface {
	DeleteRecipe(int)
}

// Repository defines methods to be used on the repository
type Repository interface {
	// AddRecipe saves a given recipe to the repository.
	DeleteRecipe(int)
}

type service struct {
	rR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) DeleteRecipe(ID int) {
	s.rR.DeleteRecipe(ID)
}
