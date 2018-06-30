package http

import (
	"net/http"

	"github.com/hoflish/url-shortener/api/urlshorten"

	models "github.com/hoflish/url-shortener/api/models"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HTTPURLShortenHandler struct {
	USUsecase urlshorten.URLShortenUsecase
}

// Get method gets information for a specified short URL
/*
	TODO: sanitize and validate shortUrl query param
		1. [done] make sure the param is set (str != "")
		2. [done] make sure the param is a url
		3. [x] check if the url has a length equal to 'LENGTH',
			(LENGTH should be defined later using host + generated id)
		4. ...
*/
func (h *HTTPURLShortenHandler) Get(c echo.Context) error {
	qparam := "shortUrl"
	query := c.QueryParams()

	if _, ok := query[qparam]; !ok {
		return c.JSON(
			http.StatusUnprocessableEntity,
			ResponseError{Message: models.ErrorMissingQueryParam.Error()},
		)
	}

	urlShort := query.Get(qparam)

	if !urlshorten.IsRequestURL(urlShort) {
		return c.JSON(
			http.StatusBadRequest,
			ResponseError{Message: models.ErrorInvalidURL.Error()},
		)
	}

	ctx := c.Request().Context()

	item, err := h.USUsecase.Fetch(ctx, urlShort)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, item)
}

// Insert creates new Short URL
func (h *HTTPURLShortenHandler) Insert(c echo.Context) error {
	var urlShorten models.URLShorten
	err := c.Bind(&urlShorten)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if !urlshorten.IsRequestURL(urlShorten.LongURL) {
		return c.JSON(
			http.StatusBadRequest,
			ResponseError{Message: models.ErrorInvalidURL.Error()},
		)
	}

	ctx := c.Request().Context()
	res, err := h.USUsecase.Store(ctx, &urlShorten)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)

}

// NewHttpURLShortenHandler defines API endpoints
func NewHTTPURLShortenHandler(e *echo.Echo, u urlshorten.URLShortenUsecase) {
	handler := &HTTPURLShortenHandler{
		USUsecase: u,
	}
	e.GET("/api/url", handler.Get)
	e.POST("/api/url", handler.Insert)
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}
	// REVIEW: Refactor this code
	logrus.Error(err)

	switch err {
	case models.ErrorInternalServer:
		return http.StatusInternalServerError
	case models.ErrorNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
