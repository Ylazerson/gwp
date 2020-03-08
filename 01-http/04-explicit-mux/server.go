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
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	server := http.Server{
		Addr:    "127.0.0.1:8090",
		Handler: mux,
	}

	server.ListenAndServe()
}
