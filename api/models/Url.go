package models

import "time"

// Url Resource
type Url struct {
	ID        string    `json:"id"`
	LongUrl   string    `json:"long_url"`
	ShortUrl  string    `json:"short_url"`
	UrlCode   string    `json:"url_code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
