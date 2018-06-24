package http

import (
	"context"
	"net/http"

	models "github.com/hoflish/url-shortener/api/models"
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
	code := c.Param("code")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	shortUrl, err := h.UrlShortenUC.Fetch(ctx, code)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, shortUrl)
}

// NewURLShortenHttpHandler defines API endpoints
func NewURLShortenHttpHandler(e *echo.Echo, u urlshorten.URLShortenUsecase) {
	handler := &HttpURLShortenHandler{
		UrlShortenUC: u,
	}
	e.GET("/api/item/:code", handler.FetchURL)
	//e.POST("/api/item", handler.Store)

}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}

	//logrus.Error(err)
	switch err {
	case models.INTERNAL_SERVER_ERROR:

		return http.StatusInternalServerError
	case models.NOT_FOUND_ERROR:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
