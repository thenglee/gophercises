package main

import (
	"fmt"
	"os"
	// "github.com/thenglee/4-html-link-parser/parse"
)

func main() {
	filename := "ex1.html"
	f, err := os.Open(filename)
	_ = f
	if err != nil {
		panic(err)
	}

	fmt.Println("File opened successfully")
}
