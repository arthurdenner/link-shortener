package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/arthurdenner/shortener/url"
	"github.com/arthurdenner/shortener/utils"
)

// Viewer takes a writer, a response and returns
// the statistics for the URL provided in the request.
func (handler *Handlers) Viewer(w http.ResponseWriter, r *http.Request) {
	utils.SearchURLAndExecute(w, r, func(uri *url.Url) {
		json, err := json.Marshal(uri.Stats())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		utils.ReplyWithJSON(w, string(json))
	})
}
