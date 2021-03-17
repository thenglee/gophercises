package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Story map[string]StoryDetails

type StoryDetails struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func parseJson(filename string) Story {
	// open json file
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	// read jsonfile as a byte array
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	// Convert json byte array data into map
	var story Story
	err = json.Unmarshal(byteValue, &story)
	if err != nil {
		panic(err)
	}

	return story
}

type storyHandler struct {
	story Story
}

var tmpl = template.Must(template.ParseFiles("story.html"))

func (sh *storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// retrieve story title
	title := r.URL.Path[1:]
	if title == "" {
		title = "intro"
	}

	if story, ok := sh.story[title]; ok {
		err := tmpl.Execute(w, story)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	http.NotFound(w, r)
}

func main() {
	story := parseJson("gopher.json")

	http.Handle("/", &storyHandler{story})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
