package parse

import (
	"io"
	"log"
	"strings"

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
				// call recursive dfs function to collate the text in all the children TextNodes
				var text []string
				text = getLinkText(n, text)
				links = append(links, Link{Href: a.Val, Text: strings.Join(text, " ")})
				break
			}
		}
	} else {
		// No need to traverse children if current node is an <a> element node
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			links = getLinks(c, links)
		}
	}

	return links
}

func getLinkText(n *html.Node, s []string) []string {
	if n.Type == html.TextNode {
		// TODO: Trim spaces, nextlines
		s = append(s, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s = getLinkText(c, s)
	}
	return s
}
