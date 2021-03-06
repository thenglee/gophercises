package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/thenglee/5-sitemap/link"
	"github.com/thenglee/5-sitemap/sitemap"
)

func main() {
	urlPtr := flag.String("URL", "https://www.calhoun.io/", "URL to generate sitemap")
	depthPtr := flag.Int("Depth", 3, "Max number of links to get to each page")
	flag.Parse()

	// visit and retrieve links based on depth
	result := bfsTraverse(*urlPtr, *depthPtr)

	// Generate a slice of sitemap Url structs
	links := make([]*sitemap.Url, len(result))
	i := 0
	for k := range result {
		links[i] = &sitemap.Url{Loc: k}
		i++
	}

	// retrieve the generated XML based on the Url structs
	output, err := sitemap.SiteMap(links)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))

}

func bfsTraverse(rootUrl string, depth int) map[string]bool {
	queue := []string{
		rootUrl,
	}
	level := map[string]int{
		rootUrl: 1,
	}
	visited := map[string]bool{
		rootUrl: true,
	}

	i := 0

	for len(queue) > 0 {
		// get the first item from the queue
		x := queue[i]

		// break when max number of links visited
		if level[x]+1 > depth {
			break
		}

		// retrieve the links present on the page
		sameDomainLinks, err := getChildrenLinks(x)
		if err != nil {
			log.Fatal(err)
			break
		}

		// loop through each of the links
		for _, link := range sameDomainLinks {

			// If link is visited, next
			if _, ok := visited[link]; ok {
				continue
			}

			// enqueue link
			queue = append(queue, link)

			// level of link is level of x + 1
			level[link] = level[x] + 1

			// mark link as visited
			visited[link] = true
		}
		i++
	}

	return visited
}

func getChildrenLinks(url string) ([]string, error) {
	// visit url
	body, err := getHTMLBody(url)
	if err != nil {
		return nil, err
	}

	// get links in body
	links, err := getLinks(body)
	if err != nil {
		return nil, err
	}

	// retrieve same domain links
	sameDomainLinks := linksInSameDomain(url, links)
	return sameDomainLinks, nil
}

func linksInSameDomain(rootUrl string, links []link.Link) []string {
	rootUrl = trimRootUrl(rootUrl)
	var sameDomainLinks []string
	rootUrlLen := len(rootUrl)

	// Acceptable links: "rootUrl/path-with-domain", "/just-the-path"
	for _, link := range links {
		href := link.Href
		hrefLen := len(href)
		// Exclude "/" root path
		if hrefLen > 1 {
			if hrefLen > rootUrlLen {
				// Starts with same domain
				if link.Href[0:rootUrlLen] == rootUrl {
					sameDomainLinks = append(sameDomainLinks, href)
				}
				// Starts with "/"
			} else if href[0:1] == "/" {
				sameDomainLinks = append(sameDomainLinks, rootUrl+href)
			}
		}
	}
	return sameDomainLinks
}

func trimRootUrl(rootUrl string) string {
	// Remove trailing slash from rootUrl if present
	length := len(rootUrl)
	if rootUrl[length-1:length] == "/" {
		rootUrl = rootUrl[:length-1]
	}
	return rootUrl
}

func getHTMLBody(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getLinks(body []byte) ([]link.Link, error) {
	r := strings.NewReader(string(body))
	links, err := link.Parse(r)
	if err != nil {
		return nil, err
	}
	return links, nil
}
