package httphandler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"

	"github.com/hoflish/url-shortener/api/models"
	urlsh "github.com/hoflish/url-shortener/api/urlshorten"
)

type HTTPURLShortenHandler struct {
	USUsecase urlsh.URLShortenUsecase
}

// Get method gets information for a specified short URL
func (h *HTTPURLShortenHandler) Get(c *gin.Context) {
	qValue, ok := c.GetQuery("shortUrl")
	if !ok {
		e := NewAPIError(400, CodeMissingParam, errors.New("Missing 'shortUrl' query parameter"))
		c.JSON(e.Status, e)
		return
	}

	if !urlsh.IsRequestURL(qValue) {
		e := NewAPIError(422, CodeInvalidParam, fmt.Errorf("Invalid shortUrl: %s", qValue))
		c.JSON(e.Status, e)
		return
	}

	item, err := h.USUsecase.Fetch(c, qValue)
	if err != nil {
		if IsDBError(err) {
			c.JSON(500, ErrAPIInternal)
			return
		}

		if IsNotFound(err) {
			c.JSON(404, ErrAPINotFound)
			return
		}

		logrus.Error(err)
		c.JSON(520, ErrAPIUnknown)
		return
	}

	jsonData(c, http.StatusOK, item)
}

// Insert creates new Short URL
func (h *HTTPURLShortenHandler) Insert(c *gin.Context) {
	var urlSh models.URLShorten
	if err := c.ShouldBindWith(&urlSh, binding.FormPost); err != nil || !urlsh.IsRequestURL(urlSh.LongURL) {
		if err == nil {
			e := NewAPIError(422, CodeInvalidParam, fmt.Errorf("Invalid longUrl: %s", urlSh.LongURL))
			c.JSON(e.Status, e)
			return
		}

		e := NewAPIError(400, CodeInvalidParam, fmt.Errorf(err.Error()))
		c.JSON(e.Status, e)
		return
	}

	res, err := h.USUsecase.Store(c, &urlSh)
	if err != nil {
		if IsDBError(err) {
			c.JSON(500, ErrAPIInternal)
			return
		}

		if IsNotFound(err) {
			c.JSON(404, ErrAPINotFound)
			return
		}

		logrus.Error(err)
		c.JSON(520, ErrAPIUnknown)
		return
	}

	jsonData(c, http.StatusOK, res)
}

// NewHTTPURLShortenHandler defines API endpoints
func NewHTTPURLShortenHandler(u urlsh.URLShortenUsecase) *HTTPURLShortenHandler {
	return &HTTPURLShortenHandler{
		USUsecase: u,
	}
}

func jsonData(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{
		"data": data,
	})
}
