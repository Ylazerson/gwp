package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// -- --------------------------------
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

// -- --------------------------------
func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

// -- --------------------------------
// This takes in a HandlerFunc and returns a HandlerFunc
// COOOOOOOL
func log(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()

		fmt.Println("Handler function called - " + name)

		h(w, r)
	}

}

// -- --------------------------------
func main() {

	server := http.Server{
		Addr: "127.0.0.1:8090",
	}

	http.HandleFunc("/hello", log(hello))
	http.HandleFunc("/world", log(world))

	server.ListenAndServe()
}
