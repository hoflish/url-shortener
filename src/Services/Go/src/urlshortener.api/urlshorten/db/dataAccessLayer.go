package db

import (
	"github.com/gin-gonic/gin"
	"urlshortener.api/models"
)

// DataAccessLayer provides thread-safe access to a database of urlshortens.
type DataAccessLayer interface {
	// Fetch retrieves a urlshorten metadata by its ShortURL.
	Fetch(ctx *gin.Context, shortURL string) (*models.URLShorten, error)

	// Store saves a given urlshorten.
	Store(ctx *gin.Context, urlsh *models.URLShorten) (*models.URLShorten, error)

	// Close closes the database, freeing up any available resources.
	Close()
}
