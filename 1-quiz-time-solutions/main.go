package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilenamePtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilenamePtr)
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s\n", *csvFilenamePtr)
		os.Exit(1)
	}
	_ = file
}
