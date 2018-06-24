package main

import (
	"log"
	"os"
	"time"

	httpDeliver "github.com/hoflish/url-shortener/api/url/delivery/http"
	urlRepos "github.com/hoflish/url-shortener/api/url/repository"
	urlUsecase "github.com/hoflish/url-shortener/api/url/usecase"
	"github.com/labstack/echo"
)

const (
	defaultHost = "192.168.99.100:27017"
)

func main() {
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := urlRepos.CreateSession(host)
	defer session.Close()

	if err != nil {
		// TODO: Use loggin API instead (e.g logrus)
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	e := echo.New()

	ur := urlRepos.NewMgoUrlRepository(session)
	timeoutContext := time.Duration(2) * time.Second

	uu := urlUsecase.NewUrlUsecase(ur, timeoutContext)
	httpDeliver.NewUrlHttpHandler(e, uu)

	e.Start(":8080")
}
