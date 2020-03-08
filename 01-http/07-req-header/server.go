package main

import (
	"fmt"
	"net/http"
)

// -- ----------------------------------
func headers(w http.ResponseWriter, r *http.Request) {

	headerMap := r.Header

	fmt.Fprintln(w, "Entire Header Map \n")
	fmt.Fprintln(w, headerMap)
	fmt.Fprintln(w, "----------------- \n")

	fmt.Fprintln(w, "Specific Header Key/Value \n")
	fmt.Fprintln(w, headerMap["Accept"])
	fmt.Fprintln(w, "----------------- \n")

	fmt.Fprintln(w, "Specific Header Key/Value \n")
	fmt.Fprintln(w, headerMap.Get("Accept"))
	fmt.Fprintln(w, "----------------- \n")
}

// -- ----------------------------------
func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/headers", headers)

	server.ListenAndServe()
}
