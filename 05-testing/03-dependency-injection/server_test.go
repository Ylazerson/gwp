// B''H

package main

// -- -----------------------------------------
/*
NOTES

Instead of passing a Post struct, this time you pass
in your FakePost struct.

The rest of the test case is no different from the
earlier one (02-httptest).

To verify that it works, shut down your database and
try running the test case. Your earlier test cases will
fail because they need to interact with a database, but
with test doubles, an actual database is no longer needed.

---

If handleGet works properly the test case will pass; otherwise
it’ll fail. Note that the test case doesn’t actually test
the fetch method in the Post struct, which requires the
setting up (and possibly tearing down) of at least the posts table.

You don’t want to do so repeatedly, which would take a long time.

Also, you want to isolate the different parts of the web
service and test them independently to make sure they work
and that you understand what went wrong.
*/
// -- -----------------------------------------

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetPost(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Errorf("Cannot retrieve JSON post")
	}
}

func TestPutPost(t *testing.T) {
	mux := http.NewServeMux()
	post := &FakePost{}
	mux.HandleFunc("/post/", handleRequest(post))

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content":"This aint gonna hit the db","author":"Sau Sheong"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Error("Response code is %v", writer.Code)
	}

	if post.Content != "This aint gonna hit the db" {
		t.Error("Content is not correct", post.Content)
	}
}
