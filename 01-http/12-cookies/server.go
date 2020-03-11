package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

// -- ---------------------------------
func setCookie(w http.ResponseWriter, r *http.Request) {

	// -- -----------------------------
	c1Message := []byte("Go Web Programming")

	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    base64.URLEncoding.EncodeToString(c1Message),
		HttpOnly: true,
	}

	// -- -----------------------------
	c2Message := []byte("Manning Publications Co")

	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    base64.URLEncoding.EncodeToString(c2Message),
		HttpOnly: true,
	}

	// -- -----------------------------
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

// -- ---------------------------------
func getCookie(w http.ResponseWriter, r *http.Request) {

	// -- -----------------------------
	// Get a single cookie:
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}

	// -- -----------------------------
	// Get all cookies:
	cs := r.Cookies()

	// -- -----------------------------
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

// -- ---------------------------------
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
