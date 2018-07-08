package main

import (
	"os"
	"time"

	db "github.com/hoflish/url-shortener/api/urlshorten/db"
	httpDeliver "github.com/hoflish/url-shortener/api/urlshorten/delivery/http"
	usecase "github.com/hoflish/url-shortener/api/urlshorten/usecase"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

const (
	defaultHost = "192.168.99.100:27017"
)

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	dbConn, err := mgo.Dial(host)
	if err != nil {
		logrus.Panicf("Init DB: %v", err)
	}

	dbConn.SetMode(mgo.Monotonic, true)
	c := dbConn.DB("url-shortener").C("urlshorten")

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

	err = dbConn.Ping()
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
	defer dbConn.Close()

	e := echo.New()

	urlshDB := db.NewMongoDB(dbConn)
	timeoutContext := time.Duration(2) * time.Second

	uu := usecase.NewURLShortenUsecase(urlshDB, timeoutContext)
	httpDeliver.NewHTTPURLShortenHandler(e, uu)

	e.Start(":8080")
}
