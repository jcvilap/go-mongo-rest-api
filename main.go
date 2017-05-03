package main

import (
	"net/http"
	"route"
	"gopkg.in/mgo.v2"
)

func main() {
	// Connect to Mongo
	s, dbError := mgo.Dial("mongodb://localhost")

	// Create routes
	route.Route(s)

	// Start server
	serviceError := http.ListenAndServe(":8080", nil)
	
	// Log any error
	panic(dbError || serviceError)
}