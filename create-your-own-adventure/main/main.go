package main

import (
	"io/ioutil"
	"encoding/json"
	"log"
)

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

func parseJson(filename string) {
	jsonStr, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
	}

	var story Story
	json.Unmarshal(jsonStr, &story)

	log.Println(story)
}

func main() {
	parseJson("story.json")
}