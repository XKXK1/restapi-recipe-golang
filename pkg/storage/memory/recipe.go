package memory

import "time"

// Recipe defines the properties of a recipe to be listed
type Recipe struct {
	ID          int
	MealType    string
	Name        string
	Ingredients []string
	Preparation string
	Created     time.Time
}
