package urlshortener

import (
	"net/http"
	"log"
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
		log.Printf("%s: %s\n", r.Method, r.URL)

		redirectPath, ok := pathsToUrls[r.URL.Path]
		if !ok {
			fallback.ServeHTTP(w, r)
			return
		}
		
		http.Redirect(w, r, redirectPath, http.StatusPermanentRedirect)
	}
}