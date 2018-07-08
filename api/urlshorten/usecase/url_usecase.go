package usecase

import (
	"context"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/hoflish/url-shortener/api/models"
	dal "github.com/hoflish/url-shortener/api/urlshorten"
	"github.com/hoflish/url-shortener/api/urlshorten/delivery/http"
	"github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
)

// TODO: refactor this code,
const shortBaseURL = "http://192.168.99.100:8080/"

type URLShortenUsecase struct {
	DB             dal.DataAccessLayer
	contextTimeout time.Duration
}

func NewURLShortenUsecase(db dal.DataAccessLayer, timeout time.Duration) dal.DataAccessLayer {
	return &URLShortenUsecase{
		DB:             db,
		contextTimeout: timeout,
	}
}

// Fetch serves data from DB layer to delivery one
func (uc *URLShortenUsecase) Fetch(c context.Context, shortURL string) (*models.URLShorten, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	item, err := uc.DB.Fetch(ctx, shortURL)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Store saves sanitized/validated inputs into DB
func (uc *URLShortenUsecase) Store(c context.Context, urlsh *models.URLShorten) (*models.URLShorten, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	// TODO: refactor this code to be more safe
	shortID, err := shortid.Generate()
	if err != nil {
		logrus.Error(err)
		return nil, http.ErrorShortID
	}

	urlsh.ID = bson.NewObjectId()
	urlsh.CreatedAt = time.Now()
	urlsh.UpdatedAt = time.Now()
	urlsh.ShortURL = shortBaseURL + shortID

	res, err := uc.DB.Store(ctx, urlsh)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Close closes DB session
func (uc *URLShortenUsecase) Close() {}
