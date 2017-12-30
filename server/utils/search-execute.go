package utils

import (
	"net/http"
	"strings"

	"github.com/arthurdenner/shortener/url"
)

func SearchURLAndExecute(w http.ResponseWriter, r *http.Request, fn func(*url.Url)) {
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	if uri := url.Search(id); uri != nil {
		fn(uri)
	} else {
		http.NotFound(w, r)
	}
}
