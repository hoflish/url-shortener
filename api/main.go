package main

import (
	"os"
	"time"

	httpDeliver "github.com/hoflish/url-shortener/api/urlshorten/delivery/http"
	urlRepos "github.com/hoflish/url-shortener/api/urlshorten/repository"
	urlUsecase "github.com/hoflish/url-shortener/api/urlshorten/usecase"
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

	e := echo.New()

	ur := urlRepos.NewMgoURLShortenRepos(session)
	timeoutContext := time.Duration(2) * time.Second

	uu := urlUsecase.NewURLShortenUsecase(ur, timeoutContext)
	httpDeliver.NewHTTPURLShortenHandler(e, uu)

	e.Start(":8080")
}
