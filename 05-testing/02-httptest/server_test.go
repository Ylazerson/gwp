package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// -- ------------------------------------
var mux *http.ServeMux
var writer *httptest.ResponseRecorder

// -- ------------------------------------
func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

// -- ------------------------------------
// Set up the global variables that are
// used in each of the test case functions.
func setUp() {

	// -- --------------------------------
	// Create multiplexer to run tests on:
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	// -- --------------------------------
	// Create the writer, which in our case
	// records the HTTP Response:
	writer = httptest.NewRecorder()
}

// -- ------------------------------------
// No need to clean up after these tests ...
func tearDown() {
}

// -- ------------------------------------
func TestHandleGet(t *testing.T) {

	// -- --------------------------------
	// Create HTTP Request to send:
	request, _ := http.NewRequest("GET", "/post/1", nil)

	// Send request to handler:
	mux.ServeHTTP(writer, request)

	// -- --------------------------------
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	// -- --------------------------------
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Errorf("Cannot retrieve JSON post")
	}
}

// -- ------------------------------------
func TestHandlePut(t *testing.T) {

	// -- --------------------------------
	json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)

	// -- --------------------------------
	request, _ := http.NewRequest("PUT", "/post/1", json)
	mux.ServeHTTP(writer, request)

	// -- --------------------------------
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
