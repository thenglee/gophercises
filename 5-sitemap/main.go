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

	fmt.Printf("%+v\n", links)
}
