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
	router := gin.New()

	env := os.Getenv("APP_ENVIRONMENT")
	switch env {
	case "production":
		file, err := os.OpenFile("log/logrus.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err == nil {
			logrus.SetOutput(file)
			logrus.SetFormatter(&logrus.JSONFormatter{})
			logrus.SetLevel(logrus.InfoLevel)

		} else {
			logrus.Info("Failed to log to file, using default stderr")
		}
		// gin mode
		gin.SetMode(gin.ReleaseMode)
		// ginrus middleware
		router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))

	case "testing":
		gin.SetMode(gin.TestMode)

	default:
		gin.SetMode(gin.DebugMode)
		router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))
	}

	v1 := router.Group("api/v1")
	{
		v1.GET("/url", h.Get)
		v1.POST("/url", h.Insert)
	}
	return router
}
