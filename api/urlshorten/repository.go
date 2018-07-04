package urlshorten

import (
	"github.com/hoflish/url-shortener/api/models"
)

// URLShortenRepos provides thread-safe access to a database of urlshortens.
type URLShortenRepos interface {
	// Fetch retrieves a urlshorten metadata by its ShortURL.
	Fetch(shortURL string) (*models.URLShorten, error)

	// Store saves a given urlshorten.
	Store(us *models.URLShorten) (*models.URLShorten, error)

	// Close closes the database, freeing up any available resources.
	Close()
}
