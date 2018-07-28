package usecase

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"gopkg.in/mgo.v2/bson"

	"urlshortener.api/models"
	"urlshortener.api/urlshorten/db"
)

type UrlService struct {
	DB db.DataAccessLayer
}

func NewUrlService(d db.DataAccessLayer) *UrlService {
	return &UrlService{
		DB: d,
	}
}

// Get returns a URLShorten resource
func (s *UrlService) Get(c *gin.Context, shortUrl string) (*models.URLShorten, error) {
	item, err := s.DB.Fetch(c, shortUrl)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Insert insert a new URLShorten resource
func (s *UrlService) Insert(c *gin.Context, urlsh *models.URLShorten) (*models.URLShorten, error) {
	shortID, err := shortid.Generate()
	if err != nil {
		return nil, err
	}

	// TODO: Use origin constant instead
	// e.g. const shortBaseUrl = "http://example.com"
	shortBaseURL := os.Getenv("WEB_SPA_ORIGIN")

	urlsh.ID = bson.NewObjectId()
	urlsh.ShortURL = shortBaseURL + "/" + shortID
	urlsh.CreatedAt = time.Now()
	urlsh.UpdatedAt = time.Now()

	res, err := s.DB.Store(c, urlsh)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Close closes DB session
func (s *UrlService) Close() {
	panic("Not implemented !")
}
