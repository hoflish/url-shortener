package urlshorten

import (
	"context"

	"github.com/hoflish/url-shortener/api/models"
)

// DataAccessLayer provides thread-safe access to a database of urlshortens.
type DataAccessLayer interface {
	// Fetch retrieves a urlshorten metadata by its ShortURL.
	Fetch(ctx context.Context, shortURL string) (*models.URLShorten, error)

	// Store saves a given urlshorten.
	Store(ctx context.Context, us *models.URLShorten) (*models.URLShorten, error)

	// Close closes the database, freeing up any available resources.
	Close()
}
