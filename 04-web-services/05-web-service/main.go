package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

// -- ---------------------------------------
type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// -- ---------------------------------------
func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

// -- ---------------------------------------
// Main handler function
/*
Note how handleRequest also takes care of any errors that
are floated up from the request, and throws a 500 status code
(StatusInternalServerError) with the error description, if there’s an error.
*/
func handleRequest(w http.ResponseWriter, r *http.Request) {

	var err error

	// -- -----------------------------------
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}

	// -- -----------------------------------
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// -- ---------------------------------------
// Retrieve a single post.
// GET /post/1
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {

	// -- -----------------------------------
	/*
		1. Extract the URL’s path
		2. Get the id using the path.Base function.
		3. The id is a string so convert it into an integer.
	*/
	id, err := strconv.Atoi(path.Base(r.URL.Path))

	if err != nil {
		return
	}

	// -- -----------------------------------
	// Get single post from DB and put into struct:
	post, err := retrieve(id)
	if err != nil {
		return
	}

	// -- -----------------------------------
	// Marshal the struct into JSON:
	output, err := json.MarshalIndent(&post, "", "\t\t")

	if err != nil {
		return
	}

	// -- -----------------------------------
	// Write JSON to the ResponseWriter:
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return
}

// -- ---------------------------------------
// Create a single post.
// POST /post/
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {

	// -- -----------------------------------
	// Create slice of bytes:
	len := r.ContentLength
	body := make([]byte, len)

	// -- -----------------------------------
	// Read request body into slice:
	r.Body.Read(body)

	// -- -----------------------------------
	// Unmarshal JSON into struct:
	var post Post
	json.Unmarshal(body, &post)

	// -- -----------------------------------
	// Create DB record:
	err = post.create()

	if err != nil {
		return
	}

	// -- -----------------------------------
	w.WriteHeader(200)
	return
}

// -- ---------------------------------------
// Update a post
// PUT /post/1
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {

	// -- -----------------------------------
	id, err := strconv.Atoi(path.Base(r.URL.Path))

	if err != nil {
		return
	}

	// -- -----------------------------------
	post, err := retrieve(id)

	if err != nil {
		return
	}

	// -- -----------------------------------
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &post)

	// -- -----------------------------------
	err = post.update()

	if err != nil {
		return
	}

	// -- -----------------------------------
	w.WriteHeader(200)
	return
}

// -- ---------------------------------------
// Delete a post
// DELETE /post/1
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {

	// -- -----------------------------------
	id, err := strconv.Atoi(path.Base(r.URL.Path))

	if err != nil {
		return
	}

	// -- -----------------------------------
	post, err := retrieve(id)

	if err != nil {
		return
	}

	// -- -----------------------------------
	err = post.delete()

	if err != nil {
		return
	}

	// -- -----------------------------------
	w.WriteHeader(200)
	return
}
