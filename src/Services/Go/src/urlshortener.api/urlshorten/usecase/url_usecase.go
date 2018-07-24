package usecase

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"urlshortener.api/models"
	dal "urlshortener.api/urlshorten"
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

// TODO: refactor this code,
const shortBaseURL = "http://localhost:3000/"

type URLShortenUsecase struct {
	DB dal.DataAccessLayer
}

func NewURLShortenUsecase(db dal.DataAccessLayer) dal.DataAccessLayer {
	return &URLShortenUsecase{
		DB: db,
	}
}

// Fetch serves data from DB layer to delivery one
func (uc *URLShortenUsecase) Fetch(c *gin.Context, shortURL string) (*models.URLShorten, error) {
	item, err := uc.DB.Fetch(c, shortURL)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Store saves sanitized/validated inputs into DB
func (uc *URLShortenUsecase) Store(c *gin.Context, urlsh *models.URLShorten) (*models.URLShorten, error) {
	shortID, err := shortid.Generate()
	if err != nil {
		panic(err)
	}

	urlsh.ID = bson.NewObjectId()
	urlsh.ShortURL = shortBaseURL + shortID
	urlsh.CreatedAt = time.Now()
	urlsh.UpdatedAt = time.Now()

	res, err := uc.DB.Store(c, urlsh)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Close closes DB session
func (uc *URLShortenUsecase) Close() {}
