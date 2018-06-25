package repository

import (
	"context"

	"github.com/hoflish/url-shortener/api/models"
	"github.com/hoflish/url-shortener/api/url"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName        = "url-shortener"
	urlCollection = "url"
)

type mgoUrlRepository struct {
	session *mgo.Session
}

// NewMgoUrlRepository creates new session
func NewMgoUrlRepository(session *mgo.Session) url.UrlRepository {
	return &mgoUrlRepository{session.Clone()}
}

// Fetch method gets a specified Url Resource
func (mg *mgoUrlRepository) Fetch(ctx context.Context, urlId string) (*models.Url, error) {
	result := models.Url{}
	// TODO: move direct operations to database to internal functions
	if err := mg.collection().Find(bson.M{"urlid": urlId}).One(&result); err != nil {
		if err == mgo.ErrNotFound {
			return nil, models.NOT_FOUND_ERROR
		}
		logrus.Error(err)
		return nil, err
	}
	return &result, nil
}

// Store method stores a new Url Resource
func (mg *mgoUrlRepository) Store(ctx context.Context, url *models.Url) (*models.Url, error) {
	panic("Not implemented yet!")
}

// collection - unexported method - returns mongodb collection
func (mg *mgoUrlRepository) collection() *mgo.Collection {
	return mg.session.DB(dbName).C(urlCollection)
}
