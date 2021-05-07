package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/thenglee/5-sitemap-solution/link"
)

/*
	1. GET the webpage
	2. parse all the links on the page
	3. build proper urls with our links
	4. filter out any links with a different domain
	5. find all the pages (BFS)
	6. print out xml
*/

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()

	/*
		/some-path
		https://gophercises/some-path
		#fragment
		mailto:jon@calhoun.io
	*/

	pages := get(*urlFlag)

	for _, p := range pages {
		fmt.Println(p)
	}
}

func get(urlStr string) []string {
	// Get webpage
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Construct base url
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	// Retrieve links on page
	links := hrefs(resp.Body, base)

	// filter links from same domain
	return filter(links, withPrefix(base))
}

func hrefs(r io.Reader, base string) []string {
	// Parse links on page
	links, _ := link.Parse(r)

	// Filter out links and prefix with base url if start with /
	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string

	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
