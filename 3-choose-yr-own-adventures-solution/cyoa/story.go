package cyoa

import (
	"encoding/json"
	"io"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func JsonStory(r io.Reader) (Story, error) {
	var story Story
	d := json.NewDecoder(r)
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Choose Your Own Adventure</title>
</head>
<body>
  <h1>{{.Title}}</h1>
  {{range .Paragraphs}}
  <p>{{.}}</p>
  {{end}}
  <ul>
    {{range .Options}}
    <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
    {{end}}
  </ul>
</body>
</html>`
