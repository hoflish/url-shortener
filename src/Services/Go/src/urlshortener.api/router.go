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

	// CORS for http://localhost:3000 origin, allowing:
	// - POST, PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Accept",
			"X-Requested-With",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		/*AllowOriginFunc: func(origin string) bool {
			return origin == "https://example.com"
		},*/
		MaxAge: 12 * time.Hour,
	}))

	v1 := router.Group("api/v1")
	{
		v1.GET("/url", h.Get)
		v1.POST("/url", h.Insert)
	}
	return router
}
