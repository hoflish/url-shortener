package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Url Resource
type Url struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	LongUrl   string        `json:"long_url"`
	UrlId     string        `json:"url_id"` // Short URL, e.g. "http://bit.ly/Cv7u".
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
