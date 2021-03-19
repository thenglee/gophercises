package main

import (
	"fmt"
	"github.com/thenglee/4-html-link-parser/parse"
	"log"
	"os"
)

func main() {
	filename := "ex1.html"
	f, err := os.Open(filename)
	_ = f
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File opened successfully")

	// s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	// parse.GetLinks(strings.NewReader(s))

	links := parse.GetLinks(f)
	fmt.Println(links)
}
