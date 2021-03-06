package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	csvFilenamePtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimitPtr := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilenamePtr)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilenamePtr))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimitPtr) * time.Second)
	correct := 0
	answerCh := make(chan string)

problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}
