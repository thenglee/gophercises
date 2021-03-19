package parse

import (
	"io"
	"log"

	"golang.org/x/net/html"
)

type Links []Link

type Link struct {
	Href string
	Text string
}

func GetLinks(r io.Reader) Links {
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	var links Links
	links = getLinks(doc, links)
	return links
}

func getLinks(n *html.Node, links Links) Links {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, Link{Href: a.Val})
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = getLinks(c, links)
	}
	return links
}
