package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/thenglee/5-sitemap/link"
)

func main() {
	urlPtr := flag.String("URL", "https://www.calhoun.io/", "URL to generate sitemap")
	flag.Parse()

	resp, err := http.Get(*urlPtr)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	r := strings.NewReader(string(body))
	links, err := link.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%+v\n", links)

	sameDomainLinks := linksInSameDomain(*urlPtr, links)
	for _, l := range sameDomainLinks {
		fmt.Println(l)
	}
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
