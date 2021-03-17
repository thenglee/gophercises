package main

import (
	"flag"
	"fmt"
)

func main() {
	filenamePtr := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *filenamePtr)
}
