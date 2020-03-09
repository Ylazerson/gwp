// B"H

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
// -- ---------------------------------
Use this approach only if you need to upload multiple files.
// -- ---------------------------------
*/

// -- ---------------------------------
func process(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(1024)

	// -- -----------------------------
	// Print form fields:
	fmt.Fprintln(w, r.MultipartForm)

	// -- -----------------------------
	// Get uploaded file:
	fileHeader := r.MultipartForm.File["file"][0]
	file, err := fileHeader.Open()

	// -- -----------------------------
	// Print uploaded file:
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
