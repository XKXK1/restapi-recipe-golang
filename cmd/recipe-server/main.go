package main

import (
	"fmt"
	"log"
	"net/http"

	"../../pkg/adding"
	"../../pkg/deleting"
	"../../pkg/http/rest"
	"../../pkg/listing"
	"../../pkg/storage/memory"
)

// StorageType defines available storage types
type Type int

func main() {

	var adder adding.Service
	var lister listing.Service
	var deleter deleting.Service

	s := new(memory.Storage)

	adder = adding.NewService(s)
	//adding SampleData
	adder.AddSampleRecipes()

	lister = listing.NewService(s)
	deleter = deleting.NewService(s)

	// set up the HTTP server
	router := rest.Handler(adder, lister, deleter)

	fmt.Println("The recipe server is online now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
