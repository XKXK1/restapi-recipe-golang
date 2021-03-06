package listing

import "time"

// Recipe defines the properties of a recipe to be added
type Recipe struct {
	ID          int       `json:"id,string"`
	MealType    string    `json:"mealtype"`
	Name        string    `json:"name"`
	Ingredients []string  `json:"ingredients"`
	Preparation string    `json:"preparation"`
	Created     time.Time `json:"created"`
}
