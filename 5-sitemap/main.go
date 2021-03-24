package main

import (
	"flag"
	"fmt"
)

func main() {
	urlPtr := flag.String("URL", "https://www.calhoun.io/", "URL to generate sitemap")
	flag.Parse()

	fmt.Println(*urlPtr)
}
