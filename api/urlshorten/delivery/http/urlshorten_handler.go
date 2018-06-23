package http

import (
	"github.com/hoflish/url-shortener/api/urlshorten"
	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HttpURLShortenHandler struct {
	UrlShortenUC urlshorten.URLShortenUsecase
}

func (h *HttpURLShortenHandler) FetchURL(c echo.Context) error {
	panic("Not implemented")
}

func NewURLShortenHttpHandler(e *echo.Echo, u urlshorten.URLShortenUsecase) {
	handler := &HttpURLShortenHandler{
		UrlShortenUC: u,
	}
	e.GET("/api/:urlcode", handler.FetchURL)
	//e.POST("/article", handler.Store)

}
