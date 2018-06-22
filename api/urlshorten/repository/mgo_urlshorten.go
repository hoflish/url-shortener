package repository

import (
	"context"

	"github.com/hoflish/url-shortener/api/models"
	"github.com/hoflish/url-shortener/api/urlshorten"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName               = "urlShortener"
	urlshortenCollection = "urlshorten"
)

type mgoURLShortenRepository struct {
	session *mgo.Session
}

// NewMgoURLShortenRepository creates new session
func NewMgoURLShortenRepository(session *mgo.Session) urlshorten.URLShortenRepository {
	return &mgoURLShortenRepository{session.Clone()}
}

func (mg *mgoURLShortenRepository) Fetch(ctx context.Context, urlCode string) (*models.URLShorten, error) {
	repo := NewMgoURLShortenRepository(mg.session)
	defer repo.Close()

	res := models.URLShorten{}
	// TODO: move direct operations to database to internal functions
	if err := mg.collection().Find(bson.M{"url_code": urlCode}).One(&res); err != nil {
		// TODO: logging conn errors (See: https://github.com/sirupsen/logrus)
		return nil, err
	}
	// TODO: retun error "NOT_FOUND_ERROR" when no data returned
	return &res, nil
}

func (mg *mgoURLShortenRepository) Store(ctx context.Context, urlShorten *models.URLShorten) (string, error) {
	panic("Not implemented")
}

// Close terminates the session
func (mg *mgoURLShortenRepository) Close() {
	mg.session.Close()
}

// collection - unexported method - returns mongodb collection
func (mg *mgoURLShortenRepository) collection() *mgo.Collection {
	return mg.session.DB(dbName).C(urlshortenCollection)
}
