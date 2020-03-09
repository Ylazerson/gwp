package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
// -- ---------------------------------
Use this approach if you need to upload a single file.

To test, use the same client.html from `09-upload-file-1` example.
// -- ---------------------------------
*/

// -- ---------------------------------
func process(w http.ResponseWriter, r *http.Request) {

	file, _, err := r.FormFile("file")

	// -- -----------------------------
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}

}

// -- ---------------------------------
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
