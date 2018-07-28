package db

import (
	"errors"
	"sync"

	"github.com/gin-gonic/gin"

	"urlshortener.api/models"
)

// Ensure MemoryDB conforms to DataAccessLayer interface.
var _ DataAccessLayer = &MemoryDB{}

// MemoryDB is a simple in-memory persistence layer for urlshortens.
type MemoryDB struct {
	mu          sync.Mutex
	urlshortens map[string]*models.URLShorten
}

// NewMemoryDB initializes a in-memory urlshortens repos
func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		urlshortens: make(map[string]*models.URLShorten),
	}
}

// Fetch retrieves URLShorten resource by its shortURL
func (db *MemoryDB) Fetch(ctx *gin.Context, shortURL string) (*models.URLShorten, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	urlsh, ok := db.urlshortens[shortURL]
	if !ok {
		return nil, errors.New("[MemDB]: Not Found")
	}
	return urlsh, nil
}

// Store adds URLShorten resource to urlshorten repos
func (db *MemoryDB) Store(ctx *gin.Context, urlsh *models.URLShorten) (*models.URLShorten, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.urlshortens[urlsh.ShortURL] = urlsh
	return urlsh, nil
}

// Close empties urlshorten repos
func (db *MemoryDB) Close() {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.urlshortens = nil
}

// Size returns the length of urlshorten repos
func (db *MemoryDB) Size() int {
	return len(db.urlshortens)
}
