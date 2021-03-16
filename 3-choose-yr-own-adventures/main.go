package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func parseJson() Story {
	// open json file
	jsonFile, err := os.Open("gopher.json")
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

func main() {
	story := parseJson()
	for k, v := range story {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("---")
	}
}
