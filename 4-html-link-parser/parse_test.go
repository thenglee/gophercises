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
