package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/arthurdenner/shortener/handlers"
	"github.com/arthurdenner/shortener/url"
	"github.com/arthurdenner/shortener/utils"
)

var (
	port     *int
	baseURL  string
	isLogsOn *bool
)

func init() {
	port = flag.Int("p", 8888, "port")
	isLogsOn = flag.Bool("l", true, "logs on/off")

	flag.Parse()

	baseURL = fmt.Sprintf("http://localhost:%d", *port)
}

func main() {
	url.ConfigRepository(url.NewMemoryRepository())

	stats := make(chan string)
	defer close(stats)
	go registerStatistics(stats)

	hand := &handlers.Handlers{BaseURL: baseURL, IsLogsOn: isLogsOn}

	http.HandleFunc("/api/short", hand.Shortener)
	http.HandleFunc("/api/short/", hand.Shortener)
	http.HandleFunc("/api/stats/", hand.Viewer)
	http.Handle("/r/", &handlers.Redirect{Stats: stats})

	utils.Logger(isLogsOn, "Listening in the port %d...", *port)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", *port), nil,
	))
}

func registerStatistics(ids <-chan string) {
	for id := range ids {
		url.RegisterClick(id)

		utils.Logger(isLogsOn, "Click registered to %s.", id)
	}
}
