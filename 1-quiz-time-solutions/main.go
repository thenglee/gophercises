package main

import (
	"flag"
	"fmt"
	"os"
)

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	csvFilenamePtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilenamePtr)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilenamePtr))
	}
	_ = file
}
