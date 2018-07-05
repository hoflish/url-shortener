package main

import (
	"os"
	"time"

	db "github.com/hoflish/url-shortener/api/urlshorten/db"
	httpDeliver "github.com/hoflish/url-shortener/api/urlshorten/delivery/http"
	usecase "github.com/hoflish/url-shortener/api/urlshorten/usecase"
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

	conn, err := db.NewMongoDB(host)
	if err != nil {
		logrus.Panicf("Init DB: %v", err)
	}
	defer conn.Close()

	timeoutContext := time.Duration(2) * time.Second

	e := echo.New()
	uu := usecase.NewURLShortenUsecase(conn, timeoutContext)
	httpDeliver.NewHTTPURLShortenHandler(e, uu)

	e.Start(":8080")
}
