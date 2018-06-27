package http

import (
	"net/http"

	models "github.com/hoflish/url-shortener/api/models"
	"github.com/hoflish/url-shortener/api/url"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HttpUrlHandler struct {
	UUsecase url.UrlUsecase
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
func (h *HttpUrlHandler) Get(c echo.Context) error {
	qparam := "shortUrl"
	query := c.QueryParams()
	
	if _, ok := query[qparam]; !ok {
		return c.JSON(
			http.StatusUnprocessableEntity,
			ResponseError{Message: models.MISSING_QUERY_PARAM.Error()},
		)
	}

	urlId := query.Get(qparam)

	if !url.IsRequestURL(urlId) {
		return c.JSON(
			http.StatusBadRequest,
			ResponseError{Message: models.INVALID_URL.Error()},
		)
	}

	ctx := c.Request().Context()

	item, err := h.UUsecase.Fetch(ctx, urlId)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, item)
}

// NewUrlHttpHandler defines API endpoints
func NewUrlHttpHandler(e *echo.Echo, u url.UrlUsecase) {
	handler := &HttpUrlHandler{
		UUsecase: u,
	}
	e.GET("/api/url", handler.Get)
	//e.POST("/api/item", handler.Create
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case models.INTERNAL_SERVER_ERROR:
		return http.StatusInternalServerError
	case models.NOT_FOUND_ERROR:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
