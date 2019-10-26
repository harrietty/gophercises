package main

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"net/http"
	"html/template"
)

var tpl = template.Must(template.ParseFiles("./assets/index.html"))

var storyData Story = parseJson("story.json")

type Story map[string]Arc

type Option struct {
	Text string `json:"text"`
	Arc string `json:"arc"`
}

type Arc struct {
	Title string `json:"title"`
	Paragraphs []string `json:"story"`
	Options []Option `json:"options"`
}

func parseJson(filename string) Story {
	jsonStr, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
	}

	var story Story
	json.Unmarshal(jsonStr, &story)

	return story
}

func main() {
	mux := http.NewServeMux()

	// Handle static assets
	staticFileDirectory := http.Dir("./assets/")
	fs := http.FileServer(staticFileDirectory)
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/arcs/", arcHandler)

	log.Printf("Running on port 8080\n")
	http.ListenAndServe(":8080", mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)
	tpl.Execute(w, storyData["intro"])
}

func arcHandler(w http.ResponseWriter, r *http.Request) {
	arc := r.URL.Path[len("/arcs/"):]
	tpl.Execute(w, storyData[arc])
}