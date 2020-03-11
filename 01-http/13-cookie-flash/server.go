package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

/*
// -- ------------------------------
There are many ways to implement flash messages, but one
of the most common is to store them in session cookies
that are removed when the page is refreshed.

Run:
    http://localhost:8080/set_message
	http://localhost:8080/show_message
	http://localhost:8080/show_message
	http://localhost:8080/show_message
// -- ------------------------------
*/

// -- ------------------------------
func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World!")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

// -- ------------------------------
func showMessage(w http.ResponseWriter, r *http.Request) {

	// -- --------------------------
	// Get "flash" cookie:
	c, err := r.Cookie("flash")

	// -- --------------------------
	if err != nil {

		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}

	} else {

		// -- ----------------------
		/*
			Here you’re replacing the existing cookie, essentially removing
			it altogether because the MaxAge field is a negative number and
			the Expires field is in the past.
		*/
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}

		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))

	}

}

// -- ------------------------------
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	server.ListenAndServe()
}
