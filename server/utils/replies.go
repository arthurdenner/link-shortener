package utils

import (
	"fmt"
	"net/http"

	"github.com/arthurdenner/shortener/types"
)

// ReplyWith writes down the headers to the response
func ReplyWith(w http.ResponseWriter, status int, headers types.Headers) {
	for key, value := range headers {
		w.Header().Set(key, value)
	}

	w.WriteHeader(status)
}

// ReplyWithJSON writes down the headers to the response
// with the Content-Type to application/json
func ReplyWithJSON(w http.ResponseWriter, reply string) {
	ReplyWith(w, http.StatusOK, types.Headers{
		"Content-Type": "application/json",
	})

	fmt.Fprintf(w, reply)
}
