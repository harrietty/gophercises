package main

import (
	"encoding/csv"
	"os"
	"fmt"
	"flag"
	"time"
)

func main() {
	fname := flag.String("f", "quiz.csv", "The filename of the quiz CSV file you wish to use")
	timeLimit := flag.Int("t", 30, "The time limit for the quiz in seconds (default of 30 seconds)")
	flag.Parse()
	file, _ := os.Open(*fname)
	r := csv.NewReader(file)
	questions, _ := r.ReadAll()
	score := 0

	time.AfterFunc(time.Duration(*timeLimit) * time.Second, func () {
		fmt.Println("Time up, sorry!")
		fmt.Printf("Your total score is %d out of %d\n", score, len(questions))
		os.Exit(1)
	})

	for _, line := range questions {
		q := line[0]
		a := line[1]
		fmt.Println(q)
		var userResponse string
		fmt.Scan(&userResponse)
		if userResponse == a {
			score++
		}
	}

	fmt.Printf("Your total score is %d out of %d\n", score, len(questions))
}