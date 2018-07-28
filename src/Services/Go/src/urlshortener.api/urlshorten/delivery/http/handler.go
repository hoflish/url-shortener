package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"

	"urlshortener.api/models"
	utils "urlshortener.api/urlshorten"
	"urlshortener.api/urlshorten/usecase"
)

type UrlHandler struct {
	service *usecase.UrlService
}

func NewUrlHandler(s *usecase.UrlService) *UrlHandler {
	return &UrlHandler{
		service: s,
	}
}

// Get handles a HTTP Get request to retrieve a UrlShorten resource
func (h *UrlHandler) Get(c *gin.Context) {
	qValue, ok := c.GetQuery("shortUrl")
	if !ok {
		e := NewAPIError(400, CodeMissingParam, errors.New("Missing 'shortUrl' query parameter"))
		c.JSON(e.Status, e)
		return
	}

	if !utils.IsRequestURL(qValue) {
		e := NewAPIError(422, CodeInvalidParam, fmt.Errorf("Invalid shortUrl: %s", qValue))
		c.JSON(e.Status, e)
		return
	}

	item, err := h.service.Get(c, qValue)
	if err != nil {
		// TODO: send errors to logging service
		if IsNotFound(err) {
			c.JSON(404, ErrAPINotFound)
			return
		}

		logrus.Error(err)
		c.JSON(500, ErrAPIInternal)
		return
	}

	jsonData(c, http.StatusOK, item)
}

// Insert handles a HTTP post request to create a new short URL
func (h *UrlHandler) Insert(c *gin.Context) {
	var urlSh models.URLShorten

	err := c.ShouldBindWith(&urlSh, binding.FormPost)
	if err != nil && !utils.IsRequestURL(urlSh.LongURL) {
		e := NewAPIError(400, CodeInvalidParam, fmt.Errorf("Invalid longUrl: %s", urlSh.LongURL))
		c.JSON(e.Status, e)
		return
	}

	res, err := h.service.Insert(c, &urlSh)
	if err != nil {
		if IsNotFound(err) {
			c.JSON(404, ErrAPINotFound)
			return
		}

		logrus.Error(err)
		c.JSON(500, ErrAPIInternal)
		return
	}

	jsonData(c, http.StatusCreated, res)
}

func jsonData(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{
		"data": data,
	})
}
