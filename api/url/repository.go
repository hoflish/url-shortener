package url

import (
	"context"

	"github.com/hoflish/url-shortener/api/models"
)

// UrlRepository defines methods which must be implemented by DB Driver
type UrlRepository interface {
	Fetch(ctx context.Context, urlCode string) (*models.Url, error)
	Store(ctx context.Context, url *models.Url) (*models.Url, error)
	Close()
}
