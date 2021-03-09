package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// notify done chan when timer fires
func quizTimer(t *time.Timer, done chan bool) {
	<-t.C
	done <- true
}

// prompt question and send received answer to ansCh chan
func getAnswer(numQnsAsked int, qns string, ansCh chan string) {
	var ans string
	// Prompt question
	fmt.Printf("Question %d: %s = ", numQnsAsked, qns)
	fmt.Scan(&ans)
	ansCh <- ans
}

func main() {
	numQnsAsked := 0
	numCorrectAns := 0

	resCh := make(chan string, 1)
	done := make(chan bool, 1)

	// Setup flag for filename
	fileNamePtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' (default \"problems.csv\")")
	// Setup flag for timer
	timerPtr := flag.Int("limit", 30, "the time limier for the quiz in seconds (default 30)")
	flag.Parse()

	// Open the file
	csvfile, err := os.Open(*fileNamePtr)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	// Setup timer
	timer1 := time.NewTimer(time.Duration(*timerPtr) * time.Second)
	go quizTimer(timer1, done)

	// Iterate through the records
Loop:
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			timer1.Stop()
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		numQnsAsked++
		go getAnswer(numQnsAsked, record[0], resCh)

		// Wait for answer or timer
		select {
		case res := <-resCh:
			// Check for correct answer
			if res == record[1] {
				numCorrectAns++
			}
			continue
		case <-done:
			break Loop
		}
	}

	// Print score
	fmt.Println()
	fmt.Printf("You scored %d out of %d.\n", numCorrectAns, numQnsAsked)

}
