package utils

import (
	"errors"
	"time"

	"../types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TryUntilInsert is a recursive function that checks for the required fields
// and if everything is OK, it runs until a unique ID be generated and the URL saved
func TryUntilInsert(c *mgo.Collection, url types.URL) error {
	var err error

	if url.LongURL == "" {
		err = errors.New("longUrl is required")
	}

	if err == nil {
		err = c.Insert(&types.URL{
			ID:        bson.NewObjectId(),
			CreatedAt: time.Now(),
			ShortURL:  GenerateID(),
			LongURL:   url.LongURL,
		})

		if err != nil {
			if mgo.IsDup(err) {
				TryUntilInsert(c, url)
			}
		}
	}

	return err
}
