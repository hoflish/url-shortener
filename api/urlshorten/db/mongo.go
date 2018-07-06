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
	session *mgo.Session
	c       *mgo.Collection
}

// NewMongoDB creates a new BookDatabase backed by a given Mongo server,
func NewMongoDB(addr string) (dal.DataAccessLayer, error) {
	/*
		TODO:
		- Authenticate with given credentials.
		Signature:
			NewMongoDB(addr string, cred *mgo.Credential)
	*/
	session, err := mgo.Dial(addr)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(dbName).C(urlShortenCollection)

	index := mgo.Index{
		Key:        []string{"shorturl"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		return nil, err
	}

	mongo := &mongoDB{
		session: session.Copy(),
		c:       c,
	}
	return mongo, nil
}

// Fetch method gets a specified Url Resource
func (db *mongoDB) Fetch(ctx context.Context, shortURL string) (*models.URLShorten, error) {
	result := models.URLShorten{}

	if err := db.c.Find(bson.M{"shorturl": shortURL}).One(&result); err != nil {
		if err == mgo.ErrNotFound {
			return nil, models.ErrorNotFound
		}
		return nil, err
	}
	return &result, nil
}

// Store method stores a new Url Resource
func (db *mongoDB) Store(ctx context.Context, us *models.URLShorten) (*models.URLShorten, error) {
	if err := db.c.Insert(us); err != nil {
		return nil, err
	}
	return us, nil
}

// Close closes database connection
func (db *mongoDB) Close() {
	if db.session != nil {
		db.session.Close()
	}
}
