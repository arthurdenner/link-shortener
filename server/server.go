package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"./types"
	"./utils"
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
	session, err := mgo.Dial("localhost")
	defer session.Close()

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	utils.EnsureIndex(session)

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/urls"), allUrls(session))
	mux.HandleFunc(pat.Post("/urls"), saveURL(session))
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

func allUrls(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		c := session.DB("link-shortener").C("urls")

		var urls []types.URL

		err := c.Find(bson.M{}).All(&urls)

		if err != nil {
			utils.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed get all urls: ", err)

			return
		}

		respBody, err := json.MarshalIndent(urls, "", "  ")

		if err != nil {
			log.Fatal(err)
		}

		utils.ReplyWithJSON(w, respBody, http.StatusOK)
	}
}

func saveURL(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		var url types.URL

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&url)

		if err != nil {
			utils.ErrorWithJSON(w, "You didn't pass any data!", http.StatusBadRequest)

			return
		}

		c := session.DB("link-shortener").C("urls")

		err = utils.TryUntilInsert(c, url)

		if err != nil {
			utils.ErrorWithJSON(w, err.Error(), http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Location", r.URL.Path+"/"+url.ShortURL)
		w.WriteHeader(http.StatusCreated)
	}
}
