package urlshortener

import (
	"net/http"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	// If the map includes the URL we care about, redirect
	return func (w http.ResponseWriter, r *http.Request) {
		redirectPath, ok := pathsToUrls[r.URL.Path]
		if !ok {
			fallback.ServeHTTP(w, r)
			return
		}
		
		http.Redirect(w, r, redirectPath, http.StatusPermanentRedirect)
	}
}

type Path struct {
	Path string
	Url string
}

type Paths struct {
	Paths []Path `yaml:"paths"`
}

func parseYaml(filename string) Paths {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading %s #%v ",filename, err)
	}

	var p Paths
	err = yaml.Unmarshal(yamlFile, &p)
	if err != nil {
			log.Fatalf("Unmarshal: %v", err)
	}

	return p
}

func createMap(paths []Path) map[string]string {
	m := make(map[string]string)
	for _, val := range paths {
		m[val.Path] = val.Url
	}
	return m
}

func YamlHandler(filename string, fallback http.Handler) http.HandlerFunc {
	yamlStuct := parseYaml(filename)
	mapOfPaths := createMap(yamlStuct.Paths)

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: %s\n", r.Method, r.URL)
		redirectPath, ok := mapOfPaths[r.URL.Path]
		if !ok {
			fallback.ServeHTTP(w, r)
			return
		}

		http.Redirect(w, r, redirectPath, http.StatusPermanentRedirect)
	}
}