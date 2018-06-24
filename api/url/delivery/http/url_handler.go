package http

import (
	"context"
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

func (h *HttpUrlHandler) GetByCode(c echo.Context) error {
	code := c.Param("code")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	item, err := h.UUsecase.Fetch(ctx, code)

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
	e.GET("/api/url/:code", handler.GetByCode)
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
