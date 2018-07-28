package db

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"urlshortener.api/models"
)

type mongoDB struct {
	Sess *mgo.Session
}

// NewMongoDB creates a new DB backed by a given Mongo server,
func NewMongoDB(Sess *mgo.Session) DataAccessLayer {
	return &mongoDB{Sess}
}

// Fetch method gets a specified Url Resource
func (db *mongoDB) Fetch(ctx *gin.Context, shortURL string) (*models.URLShorten, error) {
	s := db.Sess.Clone()
	defer s.Close()

	result := models.URLShorten{}
	if err := db.collection().Find(bson.M{"shorturl": shortURL}).One(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Store method stores a new Url Resource
func (db *mongoDB) Store(ctx *gin.Context, us *models.URLShorten) (*models.URLShorten, error) {
	s := db.Sess.Clone()
	defer s.Close()

	if err := db.collection().Insert(us); err != nil {
		return nil, err
	}
	return us, nil
}

func (db *mongoDB) collection() *mgo.Collection {
	return db.Sess.DB(Name).C(Collection)
}

// Close closes database connection
func (db *mongoDB) Close() {
	db.Sess.Close()
}
