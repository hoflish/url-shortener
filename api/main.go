package main

import (
	"net/http"
	"os"
	"time"

	"github.com/hoflish/url-shortener/api/urlshorten/db"
	"github.com/hoflish/url-shortener/api/urlshorten/delivery/http"
	"github.com/hoflish/url-shortener/api/urlshorten/usecase"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

const (
	defaultHost = "127.0.0.1:27017"
)

func main() {
	// Setup original DB session
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

	DB := db.NewMongoDB(sess)
	ucs := usecase.NewURLShortenUsecase(DB)
	handler := httphandler.NewHTTPURLShortenHandler(ucs)

	// HTTP Web server
	router := setupRouter(handler)
	srv := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	srv.ListenAndServe() // listen and serve on 0.0.0.0:8080
}
