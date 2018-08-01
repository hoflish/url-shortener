package db

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"urlshortener.api/models"
)

type mongoDB struct {
	Sess *mgo.Session
}

// NewMongoDB creates a new DB backed by a given Mongo server,
func NewMongoDB(Addr string, options ...func(*mgo.Session)) DataAccessLayer {
	sess, err := mgo.Dial(Addr)
	if err != nil {
		logrus.Panicf("Init DB: %v", err)
	}
	defer sess.Close()
	
	for _, option := range options {
		option(sess)
	}

	err = sess.Ping()
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	return &mongoDB{Sess: sess}
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
