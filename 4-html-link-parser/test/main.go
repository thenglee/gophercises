package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		// c := n.FirstChild.FirstChild.NextSibling
		if n != nil {
			fmt.Printf("n Type: %v\n", n.Type)
			fmt.Printf("n Data: %v\n", n.Data)
			fmt.Printf("n Attr: %v\n", n.Attr)
			// fmt.Printf("n FirstChild: %v\n", n.FirstChild)
			// fmt.Printf("n NextSibling: %v\n", n.NextSibling)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
