package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../../adding"
	"../../listing"
	"github.com/gorilla/mux"
)

func Handler(a adding.Service, l listing.Service) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/recipes", getRecipes(l)).Methods("GET")
	r.HandleFunc("/recipes/{id}", getRecipe(l)).Methods("GET")
	r.HandleFunc("/recipes", addRecipe(a)).Methods("POST")

	return r
}


func addRecipe(s adding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	var newRecipe adding.Recipe

	err := json.NewDecoder(r.Body).Decode(&newRecipe)
	if err != nil {
		fmt.Printf("An error occured while decoding the json: json was %s \n", r.Body)
	}
	s.AddRecipe(newRecipe)
	json.NewEncoder(w).Encode("New Recipe added.")

	}	
}

// getRecipes returns a collection of existing recipes
func getRecipes(s listing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	list := s.GetRecipes()
	json.NewEncoder(w).Encode(list)
	}
}

// getRecipes returns an existing recipe
func getRecipe(s listing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request){

		params := mux.Vars(r)

	ID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid recipe ID, it must be a number.", params["id"]), http.StatusBadRequest)
			return
		}

	recipe := s.GetRecipe(ID)


		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(recipe)
	}
}


