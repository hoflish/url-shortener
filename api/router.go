package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hoflish/url-shortener/api/urlshorten/delivery/http"
)

// setupRouter returns a framework's instance
func setupRouter(h *httphandler.HTTPURLShortenHandler) *gin.Engine {
	r := gin.New()
	r.GET("/api/url", h.Get)
	r.POST("/api/url", h.Insert)

	return r
}
