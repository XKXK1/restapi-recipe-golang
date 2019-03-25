package main

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"../../pkg/listing"
)

func TestGetRecipe(t *testing.T) {
	want := listing.Recipe{
		MealType:    "Dinner",
		Name:        "Mac&Cheese",
		Ingredients: []string{"200g Maccharoni", "200g Cheese"},
		Preparation: "Cook Maccharoni and add Cheese.",
	}

	var got listing.Recipe

	response, err := http.Get("http://0.0.0.0:8080/recipes/1")
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

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
