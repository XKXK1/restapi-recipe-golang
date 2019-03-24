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

func main() {

	var adder adding.Service
	var lister listing.Service
	var deleter deleting.Service

	// create memory where the data will be stored
	s := new(memory.Storage)

	//define services and pass storage as argument
	adder = adding.NewService(s)
	adder.AddSampleRecipes()
	lister = listing.NewService(s)
	deleter = deleting.NewService(s)

	// set up the HTTP server and add services
	router := rest.Handler(adder, lister, deleter)

	fmt.Println("The recipe server is online now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
