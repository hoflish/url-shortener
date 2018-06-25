package main

import (
	"os"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/hoflish/url-shortener/api/models"
	httpDeliver "github.com/hoflish/url-shortener/api/url/delivery/http"
	urlRepos "github.com/hoflish/url-shortener/api/url/repository"
	urlUsecase "github.com/hoflish/url-shortener/api/url/usecase"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

const (
	defaultHost = "192.168.99.100:27017"
)

func main() {
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := urlRepos.Init(host)
	defer session.Close()

	if err != nil {
		logrus.Panicf("Init DB: %v", err)
	}

	// Feed db
	data := models.Url{
		ID:        bson.NewObjectId(),
		LongUrl:   "https://www.facebook.com/",
		UrlId:     "hof.li/C7aE",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := session.DB("url-shortener").C("url").Insert(&data); err != nil {
		logrus.Errorf("Feed DB: %v", err)
	}
	// end Feed db

	e := echo.New()

	ur := urlRepos.NewMgoUrlRepository(session)
	timeoutContext := time.Duration(2) * time.Second

	uu := urlUsecase.NewUrlUsecase(ur, timeoutContext)
	httpDeliver.NewUrlHttpHandler(e, uu)

	e.Start(":8080")
}
