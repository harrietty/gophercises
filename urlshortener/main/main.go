package main

import (
	"log"
	"net/http"
	"time"
)

// Types must be declared outside of functions
type timeHandler struct {
	format string
}

/*
 timeHandler is the receiver ("this")
 * = a pointer receiver, gives you direct access to the structure in memory

	The Handler interface has been defined as below:
	type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
	}

	Therefore, any structures with a ServeHTTP method that matches the above signature can be used as a Handler
*/
func (th timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	mux := http.NewServeMux()
	
	// HandleFunc expects a function with (w http.ResponseWriter, r *http.Request) sig
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/foo", foo)

	// Handle expects a struct that implements the handler interface
	mux.Handle("/time", timeHandler{format: time.RFC1123})

	log.Println("Starting the server on :8080")

	http.ListenAndServe(":8080", mux)
}

func hello (w http.ResponseWriter, r *http.Request) {
	log.Println("Hello world!")
	w.Write([]byte("Welcome to the / route"))
}

func foo (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the /foo route!\n"))
}