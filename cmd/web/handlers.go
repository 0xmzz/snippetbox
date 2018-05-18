package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// Use r.URL.Path to check whether the request URL path exactly matches "/".
   // If it doesn't, use the w.WriteHeader() method to send a 404 status code and
   // the w.Write() method to write a "Not Found" response body. Importantly, we
   // then return from the function so that the subsequent code is not executed.
   if r.URL.Path != "/" {
	   //w.WriteHeader(404)
	   //w.Write([]byte("Not Found"))
	   // Use the http.NotFound() function to send a 404 Not Found response.
	   http.NotFound(w, r)
	   return
   }

   w.Write([]byte("Hello from MZZ's Snippetbox"))
}
