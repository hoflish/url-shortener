package main

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"urlshortener.api/urlshorten/delivery/http"
)

// SetupRouter returns a framework's instance
func SetupRouter(h *http.UrlHandler) *gin.Engine {
	router := gin.New()

	switch env := os.Getenv("APP_ENVIRONMENT"); env {
	case "production":
		file, err := os.OpenFile("logrus.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			logrus.Info("Failed to log to file, using default stderr")
		} else {
			logrus.SetOutput(file)
			logrus.SetFormatter(&logrus.JSONFormatter{})
			logrus.SetLevel(logrus.InfoLevel)
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

	// CORS for SPA Client origin, allowing:
	// - GET, POST and OPTIONS methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours

	origin := os.Getenv("WEB_SPA_ORIGIN")
	if origin == "" {
		origin = "http://0.0.0.0" // TODO: grab origin from config
	}
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{origin},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{
			"Accept",
			"Content-Length",
			"Content-Type",
			"Origin",
			"X-Requested-With",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := router.Group("api/v1")
	{
		v1.GET("/url", h.Get)
		v1.POST("/url", h.Insert)
	}
	return router
}
