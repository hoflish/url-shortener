package urlshorten

import (
	"context"

	"github.com/hoflish/url-shortener/api/models"
)

// URLShortenUsecase defines methods which handle business logic
type URLShortenUsecase interface {
	// Fetch returns a urlshorten resource by its ShortURL.
	Fetch(ctx context.Context, shortURL string) (*models.URLShorten, error)

	// Store creates a new urlshorten.
	Store(ctx context.Context, us *models.URLShorten) (*models.URLShorten, error)

	// Close closes the database.
	Close()
}
