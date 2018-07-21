package urlshorten

import (
	"urlshortener.api/models"
	"github.com/gin-gonic/gin"
)

// URLShortenUsecase defines methods which handle business process
type URLShortenUsecase interface {
	// Fetch returns a urlshorten resource by its ShortURL.
	Fetch(ctx *gin.Context, shortURL string) (*models.URLShorten, error)

	// Store creates a new urlshorten.
	Store(ctx *gin.Context, urlsh *models.URLShorten) (*models.URLShorten, error)

	// Close closes the database.
	Close()
}
