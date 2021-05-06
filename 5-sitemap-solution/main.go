package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()

	fmt.Println(*urlFlag)

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	/*
		1. GET the webpage
		2. parse all the links on the page
		3. build proper urls with our links
		4. filter out any links with a different domain
		5. find all the pages (BFS)
		6. print out xml
	*/
}
