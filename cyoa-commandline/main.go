package main

import (
	"io/ioutil"
	"fmt"
	"bufio"
	"os"
	"strings"
	"encoding/json"
	"strconv"
)

type Option struct {
	Text string `json:"text"`
	Arc string `json:"acr"`
}

type Arc struct {
	Title string `json:"title"`
	Paragraphs []string `json:"story"`
	Options []Option `json:"options"`
}

type Story map[string]Arc

func main() {
	fmt.Println("Are you ready for an adventure? Type Y or N")

	in := bufio.NewReader(os.Stdin)

	response, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	response = strings.TrimSpace(strings.ToLower(response))

	if response != "y" {
		fmt.Println("Maybe another time!")
		return		
	}

	// parse all of the JSON
	jsonStr, err := ioutil.ReadFile("story.json")
	if err != nil {
		fmt.Printf("Error reading JSON file: %v\n", err)
	}

	var story Story
	json.Unmarshal(jsonStr , &story)
	
	fmt.Println("----------------- LET'S BEGIN! ------------------------")

	// create func getNextArcIndex which displays the text and options and returns the next option name
	getNextArc(story["intro"])
	// while the arcIndex is not 0
		// call getNextArcIndex
	

	// When arcIndex is 0 again, display a goodbye message and quit
}

func getNextArc(arc Arc) {
	for _, val := range arc.Paragraphs {
		fmt.Printf("\n")
		fmt.Println(val)
		fmt.Printf("\n")
	}
	fmt.Println("----------------- WHAT NEXT? --------------------")

	for i, val := range arc.Options {
		fmt.Println(strconv.Itoa(i + 1), ") ", val.Text)
	}

	var opts []string
	for j := 0; j < len(arc.Options); j++ {
		opts = append(opts, strconv.Itoa(j + 1))
	}

	// [1, 2, 3, 4, 5 or 6]
	strOpts := ""
	for k := 0; k < len(opts); k++ {
		if k == len(opts) - 2 {
			strOpts += opts[k] +  " or "
		} else if k == len(opts) - 1 {
			strOpts += opts[k]
		} else {
			strOpts += opts[k] + ", "
		}
	}

	fmt.Println("\n Type", strOpts, "to continue.")
}