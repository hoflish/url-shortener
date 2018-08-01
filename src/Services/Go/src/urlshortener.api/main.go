package main

import (
	"net/http"
	"os"
	"time"

	"urlshortener.api/urlshorten/db"

	"github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"

	httpDelivery "urlshortener.api/urlshorten/delivery/http"
	"urlshortener.api/urlshorten/usecase"
)

func main() {
	host, port := "0.0.0.0", "27017"
	if h := os.Getenv("DB_HOST"); h != "" {
		host = h
	}
	addr := host + ":" + port

	mode := func(sess *mgo.Session) {
		sess.SetMode(mgo.Monotonic, true)
	}

	indexing := func(sess *mgo.Session) {
		c := sess.DB(db.Name).C(db.Collection)
		index := mgo.Index{
			Key:        []string{"shorturl"},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}
		err := c.EnsureIndex(index)
		if err != nil {
			logrus.Error(err)
		}
	}

	datastore := db.NewMongoDB(addr, mode, indexing)
	usecases := usecase.NewUrlService(datastore)
	urlHandler := httpDelivery.NewUrlHandler(usecases)

	// HTTP Web server
	router := SetupRouter(urlHandler)
	srv := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	srv.ListenAndServe() // listen and serve on 0.0.0.0:8080
}
