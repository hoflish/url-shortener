package urlshorten

import (
	"context"

	"github.com/hoflish/url-shortener/api/models"
)

// URLShortenRepos defines methods which must be implemented by DB Driver
type URLShortenRepos interface {
	Fetch(ctx context.Context, shortURL string) (*models.URLShorten, error)
	Store(ctx context.Context, us *models.URLShorten) (*models.URLShorten, error)
}
