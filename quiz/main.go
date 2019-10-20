package main

import (
	"encoding/csv"
	"os"
	"fmt"
	"flag"
	"time"
	"strings"
	"bufio"
)

func main() {
	// Parse the flags
	fname := flag.String("f", "quiz.csv", "The filename of the quiz CSV file you wish to use")
	timeLimit := flag.Int("t", 30, "The time limit for the quiz in seconds (default of 30 seconds)")
	flag.Parse()

	// Read the CSV file
	file, err := os.Open(*fname)
	if err != nil {
		fmt.Printf("Cannot find file %s\n", *fname)
		os.Exit(1)
	}

	in := bufio.NewReader(os.Stdin)
	
	r := csv.NewReader(file)
	lines, _ := r.ReadAll()
	score := 0

	fmt.Println("Are you ready to play? Press enter to continue.")
	in.ReadString('\n')

	time.AfterFunc(time.Duration(*timeLimit) * time.Second, func () {
		fmt.Println("Time up, sorry!")
		fmt.Printf("Your total score is %d out of %d\n", score, len(lines))
		os.Exit(1)
	})

	problems := parseLines(lines)

	for _, p := range problems {
		fmt.Println(p.question)

		// Collect answer in a separate answer channel
		answerCh := make(chan string)
		go func() {
			answer, err := in.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}
			answerCh <- strings.TrimSpace(answer)
		}()

		select {
			case answer := <-answerCh:
				if answer == p.answer {
					score++
				}
		}
	}

	fmt.Printf("Your total score is %d out of %d\n", score, len(problems))
}

func parseLines(lines [][]string) []problem {
	result := make([]problem, len(lines))
	for i, line := range lines {
		result[i] = problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}

	return result
}

type problem struct {
	question string
	answer string
}
