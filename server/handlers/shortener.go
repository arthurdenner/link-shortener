package handlers

import (
	"fmt"
	"net/http"

	"github.com/arthurdenner/shortener/types"
	"github.com/arthurdenner/shortener/url"
	"github.com/arthurdenner/shortener/utils"
)

// Shortener takes a writer, a response and checks
// if there is a shortURL for the link provided.
// If there is, returns it, if it doesn't, creates one.
func (handler *Handlers) Shortener(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.ReplyWith(w, http.StatusMethodNotAllowed, types.Headers{
			"Allow": "POST",
		})

		return
	}

	extractedURL := utils.ExtractURL(r)

	uri, isNew, err := url.SearchOrCreateNewUrl(extractedURL)

	if err != nil {
		utils.ReplyWith(w, http.StatusBadRequest, nil)

		return
	}

	var status int
	if isNew {
		status = http.StatusCreated
	} else {
		status = http.StatusOK
	}

	shortURL := fmt.Sprintf("%s/r/%s", handler.BaseURL, uri.Id)

	utils.ReplyWith(w, status, types.Headers{
		"Link": fmt.Sprintf(
			"<%s/api/stats/%s>; rel=\"stats\"",
			handler.BaseURL, uri.Id,
		),
		"Location": shortURL,
	})

	utils.Logger(handler.IsLogsOn,
		"URL %s shortened successfully to %s.",
		uri.Destiny, shortURL,
	)
}
