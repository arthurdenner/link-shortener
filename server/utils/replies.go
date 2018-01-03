package utils

import (
	"fmt"
	"net/http"

	"../types"
)

// ReplyWith writes down the headers to the response
func ReplyWith(w http.ResponseWriter, code int, headers types.Headers) {
	for key, value := range headers {
		w.Header().Set(key, value)
	}

	w.WriteHeader(code)
}

// ReplyWithJSON writes down the headers to the response
// with the Content-Type to application/json
func ReplyWithJSON(w http.ResponseWriter, json []byte, code int) {
	ReplyWith(w, code, types.Headers{
		"Content-Type": "application/json",
	})

	w.Write(json)
}

// ErrorWithJSON writes down the headers to the response
// with the Content-Type to application/json
func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	ReplyWith(w, code, types.Headers{
		"Content-Type": "application/json",
	})

	fmt.Fprintf(w, "{ \"message\": %q} ", message)
}
