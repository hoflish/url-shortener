package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// URLShorten holds metadata about a shortened URL
type URLShorten struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	LongURL   string        `json:"long_url" form:"longUrl" binding:"required"` // Long URL, e.g. "https://stackoverflow.com/"
	ShortURL  string        `json:"short_url"`                                  // Short URL, e.g. "http://bit.ly/Cv7u"
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
