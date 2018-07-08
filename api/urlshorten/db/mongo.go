package db

import (
	"context"

	"github.com/hoflish/url-shortener/api/models"
	dal "github.com/hoflish/url-shortener/api/urlshorten"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName               = "url-shortener"
	urlShortenCollection = "urlshorten"
)

type mongoDB struct {
	Sess *mgo.Session
}

// NewMongoDB creates a new DB backed by a given Mongo server,
func NewMongoDB(Sess *mgo.Session) dal.DataAccessLayer {
	return &mongoDB{Sess: Sess.Clone()}
}

// Fetch method gets a specified Url Resource
func (db *mongoDB) Fetch(ctx context.Context, shortURL string) (*models.URLShorten, error) {
	result := models.URLShorten{}
	if err := db.collection().Find(bson.M{"shorturl": shortURL}).One(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// Store method stores a new Url Resource
func (db *mongoDB) Store(ctx context.Context, us *models.URLShorten) (*models.URLShorten, error) {
	if err := db.collection().Insert(us); err != nil {
		return nil, err
	}

	return us, nil
}

func (db *mongoDB) collection() *mgo.Collection {
	return db.Sess.DB(dbName).C(urlShortenCollection)
}

// Close closes database connection
func (db *mongoDB) Close() {
	db.Sess.Close()
}
