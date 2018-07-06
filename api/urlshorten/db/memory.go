package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/hoflish/url-shortener/api/models"

	dal "github.com/hoflish/url-shortener/api/urlshorten"
)

// Ensure MemoryDB conforms to DataAccessLayer interface.
var _ dal.DataAccessLayer = &MemoryDB{}

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

// Close empties urlshorten repos
func (db *MemoryDB) Close() {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.urlshortens = nil
}

// Fetch retrieves URLShorten resource by its shortURL
func (db *MemoryDB) Fetch(ctx context.Context, shortURL string) (*models.URLShorten, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	urlsh, ok := db.urlshortens[shortURL]
	if !ok {
		return nil, fmt.Errorf("Memorydb: URLShorten not found with shortURL %s", shortURL)
	}

	return urlsh, nil
}

// Store adds URLShorten resource to urlshorten repos
func (db *MemoryDB) Store(ctx context.Context, urlsh *models.URLShorten) (*models.URLShorten, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.urlshortens[urlsh.ShortURL] = urlsh

	return urlsh, nil
}

// Size returns the length of urlshorten repos
func (db *MemoryDB) Size() int {
	return len(db.urlshortens)
}