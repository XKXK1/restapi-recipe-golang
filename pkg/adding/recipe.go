package adding

// Recipe defines the properties of a recipe to be added
type Recipe struct {
	MealType    string `json:"mealtype"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
	Preparation string `json:"preparation"`
}
