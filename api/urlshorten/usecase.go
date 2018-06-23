package urlshorten

import (
	"context"

	"github.com/hoflish/url-shortener/api/models"
)

// URLShortenUsecase defines methods which must be implemented by usecase handler
type URLShortenUsecase interface {
	Fetch(ctx context.Context, urlCode string) (*models.URLShorten, error)
	Store(ctx context.Context, urlShorten *models.URLShorten) (string, error)
	Close()
}
