// (c) 2019 Christian Bargmann
//
// This project serves for teaching purposes in the CloudWP with Stefan Sarstedt
// at the University of Applied Sciences in Hamburg. The project provides a basic framework for a Restful API,
// which can be used to manage courses of study via simple web calls.
//
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Recipe represents a recipe for a meal
type Recipe struct {
	ID           string     `json:"id"`
	MealType	 string		`json: "mealtype"`
	Name         string     `json:"name"`
	Ingredients	 string		`json: "ingredients"`
	Preparation  string     `json:"preparation"`
}


// Init a slice recipes to store our data and mock a database
var recipes []Recipe

// getRecipes returns a collection of existing recipes
func getRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

// getRecipes returns an existing recipe
func getRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get the parameters out of the http request
	params := mux.Vars(r)

	// Loop through recipes and find one with the id from the params
	// I know you can do it better, guys :-)
	for _, recipe := range recipes {
		if recipe.ID == params["id"] {
			json.NewEncoder(w).Encode(recipe)
			return
		}
	}
	json.NewEncoder(w).Encode(&Recipe{})
}

// createRecipe creates a new recipe
func createRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var recipe Recipe

	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		fmt.Printf("An error occured while decoding the json: json was %s \n", r.Body)
	}

	// Mock our recipe id
	// attention: ids may not be unique using this implementation but serves well as an example
	recipe.ID = strconv.Itoa(rand.Intn(100000))
	recipes = append(recipes, recipe)
	json.NewEncoder(w).Encode(recipe)
}

/*
// updateStudiengang updates an existing studiengang
func updateStudiengang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// again, get our params out of the request
	params := mux.Vars(r)

	// Loop through all our study courses and update matching one if found
	// Again, does this implementation does not scale well :-)
	for index, studiengang := range studiengaenge {
		if studiengang.ID == params["id"] {

			// we will simply replace the found studiengang with our new spec
			studiengaenge = append(studiengaenge[:index], studiengaenge[index+1:]...)

			var stg Studiengang

			err := json.NewDecoder(r.Body).Decode(&stg)
			if err != nil {
				fmt.Printf("An error occured while decoding the json: json was %s \n", r.Body)
			}

			stg.ID = params["id"]
			studiengaenge = append(studiengaenge, stg)
			json.NewEncoder(w).Encode(stg)
			return
		}
	}
}
*/

// deleteRecipe deletes an existing recipe
func deleteRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range recipes {
		if item.ID == params["id"] {
			recipes = append(recipes[:index], recipes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(recipes)
}

// getrandomRecipe returns a random recipe of requested type
func getRandomRecipe(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var outRecipe Recipe
	params := mux.Vars(r)

	//quick but dirty way to find a random recipe of given type
	for outRecipe.ID == "" {
		ranNumber := rand.Intn(len(recipes))
		if recipes[ranNumber].MealType == params["mealtype"]{
			outRecipe = recipes[ranNumber]
			break
		}
	}
	json.NewEncoder(w).Encode(outRecipe)
}

// main launches our simple studiengang restful api
// First we create some sample data. Thereafter we define our api routes and corresponding functions.
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	recipes = append(recipes, Recipe{ID: "1", MealType: "Breakfast", Name: "Oatmeal with cinnamon", Ingredients: "50g Oats, 100ml water, 1tsp cinnamon", Preparation: "Cook oats with water and add cinnamon." })
	recipes = append(recipes, Recipe{ID: "2", MealType: "Breakfast", Name: "Bread with butter", Ingredients: "2 slices of Bread, 20g butter", Preparation: "Spread the butter on the bread." })
	recipes = append(recipes, Recipe{ID: "3", MealType: "Dinner", Name: "Asian Curry", Ingredients: "100g Rice, 20g curry, 2 carrots", Preparation: "Chop carrots. Cook rice with curry and carrots." })

	// Route handles & endpoints
	setupEndpoints(r)

	// Start server
	fmt.Printf("webserver is running at 0.0.0.0:8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func setupEndpoints(r  *mux.Router) {

	r.HandleFunc("/recipes", getRecipes).Methods("GET")
	r.HandleFunc("/recipes/{id}", getRecipe).Methods("GET")
	r.HandleFunc("/recipes", createRecipe).Methods("POST")
	r.HandleFunc("/recipes/{id}", deleteRecipe).Methods("DELETE")
	r.HandleFunc("/recipes/random/{mealtype}", getRandomRecipe).Methods("GET")
}
