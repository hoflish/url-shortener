package repository

import (
	"context"

	"github.com/hoflish/url-shortener/api/models"
	"github.com/hoflish/url-shortener/api/urlshorten"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName        = "url-shortener"
	urlCollection = "urlshorten"
)

type mgoURLShortenRepos struct {
	session *mgo.Session
}

// NewMgoURLShortenRepos creates new session
func NewMgoURLShortenRepos(session *mgo.Session) urlshorten.URLShortenRepos {
	return &mgoURLShortenRepos{session.Clone()}
}

// Fetch method gets a specified Url Resource
func (mg *mgoURLShortenRepos) Fetch(ctx context.Context, shortUrl string) (*models.URLShorten, error) {
	result := models.URLShorten{}
	// REVIEW: move direct operations to database to internal functions
	if err := mg.collection().Find(bson.M{"shorturl": shortUrl}).One(&result); err != nil {
		if err == mgo.ErrNotFound {
			return nil, models.ErrorNotFound
		}
		logrus.Error(err)
		return nil, err
	}
	return &result, nil
}

// Store method stores a new Url Resource
/*func (mg *mgoUrlRepository) Store(ctx context.Context, url *models.Url) (*models.Url, error) {
	panic("Not implemented yet!")
}*/

// collection - unexported method - returns mongodb collection
func (mg *mgoURLShortenRepos) collection() *mgo.Collection {
	return mg.session.DB(dbName).C(urlCollection)
}
