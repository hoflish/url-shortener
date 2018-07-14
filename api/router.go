package main

import (
	"os"
	"time"

	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/hoflish/url-shortener/api/urlshorten/delivery/http"
	"github.com/sirupsen/logrus"
)

// SetupRouter returns a framework's instance
func SetupRouter(h *httphandler.HTTPURLShortenHandler) *gin.Engine {
	env := os.Getenv("APP_ENVIRONMENT")

	if env == "production" {
		file, err := os.OpenFile("log/logrus.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err == nil {
			logrus.SetOutput(file)
			logrus.SetFormatter(&logrus.JSONFormatter{})
			logrus.SetLevel(logrus.InfoLevel)

		} else {
			logrus.Info("Failed to log to file, using default stderr")
		}

		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// ginrus middleware, which:
	r.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))

	r.GET("/api/url", h.Get)
	r.POST("/api/url", h.Insert)

	return r
}
