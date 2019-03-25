package main

import (
	"../../pkg/listing"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

const baseURL = "http://0.0.0.0:8080"

// TestGetRecipen tests getting the recipe with id "1"
func TestGetRecipe(t *testing.T) {
	want := listing.Recipe{
		MealType:    "Dinner",
		Name:        "Mac&Cheese",
		Ingredients: []string{"200g Maccharoni", "200g Cheese"},
		Preparation: "Cook Maccharoni and add Cheese.",
	}

	var got listing.Recipe

	response, err := http.Get(baseURL + "/recipes/1")
	if err != nil {
		t.Fatalf("failed to get json, %s", err)
	}

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	err = json.NewDecoder(response.Body).Decode(&got)
	if err != nil {
		t.Fatalf("failed to parse json, %s", err)
	}

	if (want.MealType != got.MealType) || (want.Name != got.Name) || !reflect.DeepEqual(want.Ingredients, got.Ingredients) || (want.Preparation != got.Preparation) {
		t.Fatalf("Test failed")
	}
}

// TestGetRecipes tests getting all 5 recipes from added samlpedata
func TestGetRecipes(t *testing.T) {
	var got []listing.Recipe
	const wantedSliceLength = 5

	response, err := http.Get(baseURL + "/recipes")
	if err != nil {
		t.Fatalf("failed to get json, %s", err)
	}

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	err = json.NewDecoder(response.Body).Decode(&got)
	if err != nil {
		t.Fatalf("failed to parse json, %s", err)
	}

	if len(got) != wantedSliceLength {
		t.Fatalf("Test failed")
	}
}

// TestAddRecipe tests adding one sample recipe
func TestAddRecipe(t *testing.T) {
	payload := fmt.Sprintf(`
    {
    "mealtype": "Breakfast",
    "name": "Pancakes",
    "Ingredients": [ "150g all purpose flour",
    				 "150ml of milk"],
    "preparation": "Add all ingredients and mix. Put in Pan."
}`)

	response, err := http.Post(baseURL+"/recipes", "application/json", strings.NewReader(payload))
	if err != nil {
		t.Fatalf("failed to get json, %s", err)
	}

	checkResponseCode(t, http.StatusOK, response.StatusCode)
}

// TestGetRandomRecipe tests getting a random recipe with "Breakfast" as Mealtype
func TestGetRandomRecipe(t *testing.T) {
	var got listing.Recipe
	const wantedMealType = "Breakfast"

	response, err := http.Get(baseURL + "/recipes/random/Breakfast")
	if err != nil {
		t.Fatalf("failed to get json, %s", err)
	}

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	err = json.NewDecoder(response.Body).Decode(&got)
	if err != nil {
		t.Fatalf("failed to parse json, %s", err)
	}

	if wantedMealType != got.MealType {
		t.Fatalf("Test failed. Wanted MealType 'Breakfast' but got Mealtype %s", got.MealType)
	}
}

// TestDeleteRecipe tests deleting the recipe with id "1"
func TestDeleteRecipe(t *testing.T) {
	request, err := http.NewRequest("DELETE", baseURL+"/recipes/1", nil)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatalf("failed to get json, %s", err)
	}
	checkResponseCode(t, http.StatusOK, response.StatusCode)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
