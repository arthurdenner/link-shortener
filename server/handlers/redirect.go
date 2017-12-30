package handlers

import (
	"net/http"

	"github.com/arthurdenner/shortener/url"
	"github.com/arthurdenner/shortener/utils"
)

// Redirect does something...
type Redirect struct {
	Stats chan string
}

// Redirect takes a writer, a response and get
// the id from the URL provided in the request.
func (red *Redirect) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	utils.SearchURLAndExecute(w, r, func(uri *url.Url) {
		http.Redirect(w, r, uri.Destiny,
			http.StatusMovedPermanently)

		red.Stats <- uri.Id
	})
}
