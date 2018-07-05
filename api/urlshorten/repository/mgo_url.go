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
	dbName               = "url-shortener"
	urlShortenCollection = "urlshorten"
)

type mongoDB struct {
	session *mgo.Session
}

// NewMongoDB creates new session
func NewMongoDB(session *mgo.Session) urlshorten.URLShortenRepos {
	return &mongoDB{session.Clone()}
}

// Close closes database connection
func (db *mongoDB) Close() {
	db.session.Close()
}

// Fetch method gets a specified Url Resource
func (db *mongoDB) Fetch(ctx context.Context, shortURL string) (*models.URLShorten, error) {
	result := models.URLShorten{}

	if err := db.collection().Find(bson.M{"shorturl": shortURL}).One(&result); err != nil {
		if err == mgo.ErrNotFound {
			return nil, models.ErrorNotFound
		}
		logrus.Error(err)
		return nil, err
	}
	return &result, nil
}

// Store method stores a new Url Resource
func (db *mongoDB) Store(ctx context.Context, us *models.URLShorten) (*models.URLShorten, error) {
	us.ID = bson.NewObjectId()

	logrus.Debug("Created At: ", us.CreatedAt)
	if err := db.collection().Insert(us); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return us, nil
}

// collection - unexported method - returns mongodb collection
func (db *mongoDB) collection() *mgo.Collection {
	return db.session.DB(dbName).C(urlShortenCollection)
}
