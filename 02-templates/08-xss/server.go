package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {

	// BAD: In reality you would NOT use this!!
	w.Header().Set("X-XSS-Protection", "0")

	t, _ := template.ParseFiles("tmpl.html")

	// GOOD: the input is escaped!
	t.Execute(w, r.FormValue("comment"))

	// BAD: the input is unescaped!
	t.Execute(w, template.HTML(r.FormValue("comment")))
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)

	// Put in the form: <script>alert('Pwnd!');</script>
	http.HandleFunc("/", form)

	server.ListenAndServe()
}
