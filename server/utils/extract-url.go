package utils

import "net/http"

// ExtractURL receives a request and extracts the given URL
func ExtractURL(r *http.Request) string {
	url := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(url)

	return string(url)
}
