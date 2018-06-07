package main

import (
	"log"
	"net/http"
)

func main() {
	//mux is a router in go, it stores a maping between the URL patterns for app and the routers
	//usually one server mux for one app containign all routes
	// Use the http.NewServeMux() function to initialize a new serve mux. Then use
	// the mux.HandleFunc() method to register our Home function as the handler for
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/new", NewSnippet)


	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

