package main

/*
-- ------------------------------------
Walkthrough:

1: go run server.go

2: Open browser to http://localhost:8090/
-- ------------------------------------
*/

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(
		writer,
		"Hello World, %s!",
		request.URL.Path[1:],
	)

}

func main() {

	// -- ------------------------------
	// This registers the handler function for the given pattern in the DefaultServeMux.
	http.HandleFunc("/", handler)

	// -- ------------------------------
	/*
		func ListenAndServe(addr string, handler Handler) error

			- It listens on the TCP network `addr`
			- It then calls Serve with `handler` to handle requests on incoming connections.

		The handler is typically nil, in which case the DefaultServeMux is used.
	*/
	http.ListenAndServe(":8090", nil)

}
