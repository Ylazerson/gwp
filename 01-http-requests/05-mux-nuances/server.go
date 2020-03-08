package main

import (
	"fmt"
	"net/http"
)

// -- --------------------------------
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi I am the root index level.")
}

// -- --------------------------------
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

// -- --------------------------------
func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

// -- --------------------------------
func main() {

	// -- ----------------------------
	mux := http.NewServeMux()

	// -- ----------------------------
	/*
		Will be called for:
			http://localhost:8090/
			http://localhost:8090/wazzup
			http://localhost:8090/world/
			http://localhost:8090/world/huh
	*/
	mux.HandleFunc("/", index)

	// -- ----------------------------
	/*
		Will be called for:
			http://localhost:8090/hello/
			http://localhost:8090/hello/why/me
	*/
	mux.HandleFunc("/hello/", hello)

	// -- ----------------------------
	/*
		Will ONLY be called for:
			http://localhost:8090/world
	*/
	mux.HandleFunc("/world", world)

	// -- ----------------------------
	server := http.Server{
		Addr:    "127.0.0.1:8090",
		Handler: mux,
	}

	server.ListenAndServe()
}
