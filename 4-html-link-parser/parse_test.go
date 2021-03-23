package main

import (
	"fmt"
	"github.com/thenglee/4-html-link-parser/parse"
	"strings"
	"testing"
)

func TestGetLinks(t *testing.T) {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	links := parse.GetLinks(strings.NewReader(s))
	expectedLinks := []parse.Link{
		parse.Link{Href: "foo", Text: "Foo"},
		parse.Link{Href: "/bar/baz", Text: "BarBaz"},
	}

	for i, link := range links {
		if link != expectedLinks[i] {
			t.Error(fmt.Sprintf("Expected %v but got %v.\n", expectedLinks[i], link))
		}
	}
}

func TestGetLinksNested(t *testing.T) {
	s := `<p>Links:</p><ul><li><a href="https://www.twitter.com/joncalhoun">Check me out on twitter<i class="fa fa-twitter" aria-hidden="true"></i></a><li><a href="https://github.com/gophercises">Gophercises is on <strong>Github</strong>!</a></ul>`
	links := parse.GetLinks(strings.NewReader(s))
	expectedLinks := []parse.Link{
		parse.Link{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"},
		parse.Link{Href: "https://github.com/gophercises", Text: "Gophercises is on Github !"},
	}

	for i, link := range links {
		if link != expectedLinks[i] {
			t.Error(fmt.Sprintf("Expected %v but got %v.\n", expectedLinks[i], link))
		}
	}
}

func TestGetLinksWithComment(t *testing.T) {
	s := `<html><body><a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a></body></html>`
	links := parse.GetLinks(strings.NewReader(s))
	expectedLinks := []parse.Link{
		parse.Link{Href: "/dog-cat", Text: "dog cat"},
	}

	for i, link := range links {
		if link != expectedLinks[i] {
			t.Error(fmt.Sprintf("Expected %v but got %v.\n", expectedLinks[i], link))
		}
	}
}
