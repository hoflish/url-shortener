package repository

import (
	"context"

	"github.com/hoflish/url-shortener/api/models"
	"github.com/hoflish/url-shortener/api/urlshorten"
	"gopkg.in/mgo.v2"
)

const (
	dbName               = "urlShortener"
	urlshortenCollection = "urlshorten"
)

type mgoURLShortenRepository struct {
	session *mgo.Session
}

func NewMgoURLShortenRepository(session *mgo.Session) urlshorten.URLShortenRepository {
	return &mgoURLShortenRepository{session.Clone()}
}

func (mg *mgoURLShortenRepository) Fetch(ctx context.Context, urlCode string) (*models.URLShorten, error) {
	repo := NewMgoURLShortenRepository(mg.session)
	defer repo.Close()
	panic("Not implemented")
}

func (mg *mgoURLShortenRepository) Close() {
	mg.session.Close()
}

func (mg *mgoURLShortenRepository) collection() *mgo.Collection {
	return mg.session.DB(dbName).C(urlshortenCollection)
}
