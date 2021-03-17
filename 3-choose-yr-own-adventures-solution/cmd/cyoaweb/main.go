package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/thenglee/3-choose-yr-own-adventures-solution/cyoa"
)

func main() {
	filenamePtr := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *filenamePtr)

	f, err := os.Open(*filenamePtr)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", story)
}
