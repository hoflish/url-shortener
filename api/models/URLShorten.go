package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// URLShorten Resource
type URLShorten struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	LongURL   string        `json:"long_url"`  // Long URL, e.g. "https://stackoverflow.com/"
	ShortURL  string        `json:"short_url"` // Short URL, e.g. "http://bit.ly/Cv7u"
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
