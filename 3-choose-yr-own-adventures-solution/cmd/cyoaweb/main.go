package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	// "html/template"

	"github.com/thenglee/3-choose-yr-own-adventures-solution/cyoa"
)

func main() {
	filenamePtr := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	portPtr := flag.Int("post", 3000, "the port to start the CYOA web application on")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *filenamePtr)

	f, err := os.Open(*filenamePtr)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	// To test out setting custom template via optional arguments
	// tpl := template.Must(template.New("").Parse("hello"))
	// h := cyoa.NewHandler(story, cyoa.WithTemplate(tpl))

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *portPtr)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portPtr), h))

}
