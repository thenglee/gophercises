package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type StoryDetails struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func parseJson() {
	// open json file
	jsonFile, err := os.Open("gopher.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully open json file")
	defer jsonFile.Close()

	// read jsonfile as a byte array
	byteValue, err := ioutil.ReadAll(jsonFile)

	var gopherAdventure map[string]StoryDetails
	json.Unmarshal(byteValue, &gopherAdventure)

	for k, v := range gopherAdventure {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("---")
	}

	// read from json file
	// parse json
	// load json data in map
}

func main() {
	parseJson()
}
