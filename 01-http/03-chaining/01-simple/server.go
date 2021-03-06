package main

import (
	"fmt"
	"net/http"
)

// -- --------------------------------
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

// -- --------------------------------
// This takes in a HandlerFunc and returns a HandlerFunc
// COOOOOOOL
func log(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Handler function called")

		h(w, r)
	}

}

// -- --------------------------------
func main() {

	server := http.Server{
		Addr: "127.0.0.1:8090",
	}

	http.HandleFunc("/hello", log(hello))

	server.ListenAndServe()
}
