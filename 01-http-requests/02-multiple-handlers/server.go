package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {

	// -- ------------------------------
	/*
			Defines parameters for running an HTTP server.

			The Server struct configuration:

		    type Server struct {
		        Addr           string
		        Handler        Handler
		        ReadTimeout    time.Duration
		        WriteTimeout   time.Duration
		        MaxHeaderBytes int
		        TLSConfig      *tls.Config
		        TLSNextProto   map[string]func(*Server, *tls.Conn, Handler)
		        ConnState      func(net.Conn, ConnState)
		        ErrorLog       *log.Logger
		    }


			For details see: https://pkg.go.dev/net/http?tab=doc#Server

	*/
	server := http.Server{
		Addr: "127.0.0.1:8090",
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
