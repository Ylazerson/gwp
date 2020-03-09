package main

import (
	"fmt"
	"net/http"
)

/*
-- -------------------------------
To test:

RUN: go run server.go

OPEN:
-- -------------------------------
*/

func process(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, r.PostFormValue("name"))
	fmt.Fprintln(w, r.PostFormValue("addr"))

}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
