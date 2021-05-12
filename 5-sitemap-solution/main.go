package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
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

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Xmlns string `xml:"xmlns,attr"`
	Urls  []loc  `xml:"url"`
}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 3, "the maximum number of links deep to traverse")
	flag.Parse()

	/*
		/some-path
		https://gophercises/some-path
		#fragment
		mailto:jon@calhoun.io
	*/

	// retrieve sitemap links
	pages := bfs(*urlFlag, *maxDepth)

	toXml := urlset{
		Xmlns: xmlns,
		Urls:  make([]loc, len(pages)),
	}
	for i, page := range pages {
		toXml.Urls[i] = loc{page}
	}

	// Print xml header
	fmt.Print(xml.Header)

	// Encode and print toXml
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}
	fmt.Println()
}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{}) // value is struct rather than bool as struct is more memory efficient
	var q map[string]struct{}         // current q to loop and visit
	nq := map[string]struct{}{        // next q, for appending the current page's links
		urlStr: struct{}{},
	}

	for i := 0; i <= maxDepth; i++ {
		// assign nq to q, create and assign a new empty map to nq
		q, nq = nq, make(map[string]struct{})

		// Break if no links in q
		if len(q) == 0 {
			break
		}

		for currentUrl, _ := range q {
			// if currentUrl is visited, continue
			if _, ok := seen[currentUrl]; ok {
				continue
			}
			// mark currentUrl as seen
			seen[currentUrl] = struct{}{}
			// retrieve links on currentUrl page and add to nq
			for _, link := range get(currentUrl) {
				// if link is not seen yet, add to nq
				if _, ok := seen[link]; !ok {
					nq[link] = struct{}{}
				}
			}
		}
	}

	ret := make([]string, 0, len(seen))
	// convert map of urls to slice of urls
	for url, _ := range seen {
		ret = append(ret, url)
	}
	return ret
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
