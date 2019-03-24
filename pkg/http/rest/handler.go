package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../../adding"
	"../../deleting"
	"../../listing"
	"github.com/gorilla/mux"
)

// Handler defines all rest-endpoints to be called
func Handler(a adding.Service, l listing.Service, d deleting.Service) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/recipes", getRecipes(l)).Methods("GET")
	r.HandleFunc("/recipes/{id}", getRecipe(l)).Methods("GET")
	r.HandleFunc("/recipes/random/{mealType}", getRandomRecipe(l)).Methods("GET")
	r.HandleFunc("/recipes", addRecipe(a)).Methods("POST")
	r.HandleFunc("/recipes/{id}", deleteRecipe(d)).Methods("DELETE")

	return r
}

func addRecipe(s adding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var newRecipe adding.Recipe

		err := json.NewDecoder(r.Body).Decode(&newRecipe)
		if err != nil {
			fmt.Printf("An error occured while decoding the json: json was %s \n", r.Body)
		}
		s.AddRecipe(newRecipe)
		if err := json.NewEncoder(w).Encode("New Recipe added."); err != nil {
			panic("Couldn't Encode message")
		}

	}
}

// getRecipes returns a collection of existing recipes
func getRecipes(s listing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetRecipes()

		if err := json.NewEncoder(w).Encode(list); err != nil {
			panic("Couldn't Encode message")
		}
	}
}

// getRecipes returns an existing recipe
func getRecipe(s listing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		ID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid recipe ID, it must be a number.", params["id"]), http.StatusBadRequest)
			return
		}

		recipe, err := s.GetRecipe(ID)
		if err == listing.ErrNotFound {
			http.Error(w, "The recipe you requested does not exist.", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(recipe); err != nil {
			panic("Couldn't Encode message")
		}
	}
}

// getRecipes returns a collection of existing recipes
func getRandomRecipe(s listing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//todo errorhandling
		params := mux.Vars(r)

		recipe := s.GetRandomRecipe(params["mealType"])

		if err := json.NewEncoder(w).Encode(recipe); err != nil {
			panic("Couldn't Encode message")
		}
	}
}

func deleteRecipe(s deleting.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		ID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid recipe ID, it must be a number.", params["id"]), http.StatusBadRequest)
			return
		}

		s.DeleteRecipe(ID)

		w.Header().Set("Content-Type", "application/json")
		outMsg := fmt.Sprintf("Recipe with ID: %d has been deleted", ID)

		if err := json.NewEncoder(w).Encode(outMsg); err != nil {
			panic("Couldn't Encode message")
		}
	}
}
