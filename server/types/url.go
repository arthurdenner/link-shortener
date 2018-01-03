package types

import "gopkg.in/mgo.v2/bson"
import "time"

// URL describes how a URL will be stored in MongoDB
type URL struct {
	ID        bson.ObjectId `bson:"_id" json:"_id,omitempty"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	ShortURL  string        `bson:"shortUrl" json:"shortUrl"`
	LongURL   string        `bson:"longUrl" json:"longUrl"`
}
