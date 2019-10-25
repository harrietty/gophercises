package main

import (
	"log"
	"net/http"
	"time"
	"flag"
	"github.com/harrietty/gophercises/urlshortener"
	"github.com/harrietty/gophercises/urlshortener/sql"
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
	yamlFileName := flag.String("f", "paths.yaml", "The name of the YAML file where shortened URLS are mapped")
	shouldSeed := flag.Bool("s", false, "Whether or not the seed file should be executed")
	
	flag.Parse()
	
	if *shouldSeed {
		seed.SeedDB()
	}

	mux := http.NewServeMux()
	
	// HandleFunc expects a function with (w http.ResponseWriter, r *http.Request) sig
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/foo", foo)

	// Handle expects a struct that implements the handler interface
	mux.Handle("/time", timeHandler{format: time.RFC1123})

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"/twitter": "https://twitter.com",
	}
	mapHandler := urlshortener.MapHandler(pathsToUrls, mux)

	yamlHandler := urlshortener.YamlHandler(*yamlFileName, mapHandler)

	dbHandler := urlshortener.DbHandler(yamlHandler)

	jsonHandler := urlshortener.JsonHandler("paths.json", dbHandler)

	log.Println("Starting the server on :8080")

	// Needs to be passed something which implements ServeHTTP
	http.ListenAndServe(":8080", jsonHandler)
}

func hello (w http.ResponseWriter, r *http.Request) {
	log.Println("Hello world!")
	w.Write([]byte("Welcome to the / route"))
}

func foo (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the /foo route!\n"))
}