package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Open the file
	csvfile, err := os.Open("problems.csv")
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
		if ans == record[1] {
			numCorrectAns++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", numCorrectAns, numQnsAsked)

}
