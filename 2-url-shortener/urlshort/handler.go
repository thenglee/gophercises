package urlshort

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url := pathsToUrls[r.RequestURI]; url != "" {
			http.Redirect(w, r, url, 302)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}
	buildMap(parsedYaml)
	return nil, nil
}

type RedirectData struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func parseYaml(yml []byte) ([]RedirectData, error) {
	var redirectData []RedirectData

	err := yaml.Unmarshal(yml, &redirectData)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return redirectData, err
	}
	fmt.Println(redirectData)
	return redirectData, nil
}

func buildMap(parsedYml []RedirectData) map[string]string {
	redirectMap := make(map[string]string)
	for _, st := range parsedYml {
		redirectMap[st.Path] = st.Url
	}
	return redirectMap
}
