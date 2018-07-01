package urlshorten

import (
	"context"

	"github.com/hoflish/url-shortener/api/models"
)

// URLShortenUsecase defines methods which must be implemented by usecase handler
type URLShortenUsecase interface {
	Fetch(ctx context.Context, shortURL string) (*models.URLShorten, error)
	Store(ctx context.Context, us *models.URLShorten) (*models.URLShorten, error)
}
