package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Setup flag for filename
	fileNamePtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' (default \"problems.csv\")")
	flag.Parse()

	// Open the file
	csvfile, err := os.Open(*fileNamePtr)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	var ans string
	numQnsAsked := 0
	numCorrectAns := 0

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		numQnsAsked++
		// Prompt question
		fmt.Printf("Question %d: %s = ", numQnsAsked, record[0])
		fmt.Scan(&ans)
		// Check for correct answer
		if ans == record[1] {
			numCorrectAns++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", numCorrectAns, numQnsAsked)

}
