package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Url Resource
type Url struct {
	ID        bson.ObjectId `bson:"_id" json:"id"` // TODO: rename json id to _id
	LongUrl   string        `json:"long_url"`
	ShortUrl  string        `json:"short_url"` // TODO: rename this field to url_id
	UrlCode   string        `json:"url_code"`  // TODO: remove this field
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
