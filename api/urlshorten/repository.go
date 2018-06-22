package repository

import (
	"context"

	"github.com/hoflish/url-shortener/api/models"
)

// URLShortenRepository defines methods which must be implemented by DB Driver
type URLShortenRepository interface {
	Fetch(ctx context.Context, urlCode string) (*models.URLShorten, error)
	Store(ctx context.Context, urlShorten *models.URLShorten) (string, error)
}
