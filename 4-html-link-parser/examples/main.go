package main

import (
	"flag"
	"fmt"
	"github.com/thenglee/4-html-link-parser/parse"
	"log"
	"os"
)

func main() {
	filenamePtr := flag.String("filename", "ex1.html", "html file to parse as links")
	flag.Parse()

	f, err := os.Open(*filenamePtr)
	_ = f
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File opened successfully")

	// s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	// links := parse.GetLinks(strings.NewReader(s))

	links := parse.GetLinks(f)
	fmt.Printf("links: %+v\n", links)
}
