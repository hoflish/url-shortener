package main

import (
	"os"
	"time"

	"github.com/hoflish/url-shortener/api/urlshorten/db"
	"github.com/hoflish/url-shortener/api/urlshorten/delivery/http"
	"github.com/hoflish/url-shortener/api/urlshorten/usecase"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

const (
	defaultHost = "127.0.0.1:27017"
)

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	sess, err := mgo.Dial(host)
	if err != nil {
		logrus.Panicf("Init DB: %v", err)
	}

	sess.SetMode(mgo.Monotonic, true)
	c := sess.DB(db.Name).C(db.Collection)

	index := mgo.Index{
		Key:        []string{"shorturl"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		logrus.Error(err)
	}

	err = sess.Ping()
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
	defer sess.Close()

	e := echo.New()

	urlshDB := db.NewMongoDB(sess)
	timeoutContext := time.Duration(2) * time.Second

	ucs := usecase.NewURLShortenUsecase(urlshDB, timeoutContext)
	h := httphandler.NewHTTPURLShortenHandler(e, ucs)

	e.GET("/api/url", h.Get)
	e.POST("/api/url", h.Insert)

	e.Start(":8080")
}
