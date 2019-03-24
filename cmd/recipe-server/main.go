package main

import (
	"fmt"
	"log"
	"net/http"

	
	"../../pkg/adding"
	"../../pkg/http/rest"
	"../../pkg/listing"
	"../../pkg/storage/memory"
	

)

// StorageType defines available storage types
type Type int

func main() {


	var adder adding.Service
	var lister listing.Service

	s := new(memory.Storage)

	adder = adding.NewService(s)
	lister = listing.NewService(s)


	// set up the HTTP server
	router := rest.Handler(adder, lister)

	fmt.Println("The recipe server is online now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
