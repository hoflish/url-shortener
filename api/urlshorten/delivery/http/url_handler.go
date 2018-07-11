package httphandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/hoflish/url-shortener/api/models"
	urlsh "github.com/hoflish/url-shortener/api/urlshorten"
)

type HTTPURLShortenHandler struct {
	USUsecase urlsh.URLShortenUsecase
}

// Get method gets information for a specified short URL
func (h *HTTPURLShortenHandler) Get(c *gin.Context) {
	qParam := "shortUrl" // required query param
	qParams := []string{qParam}
	qValue, ok := c.GetQuery(qParam)

	if !ok {
		jsonErrResponse(c, ErrorMissingParam, qParams)
		return
	}

	if !urlsh.IsRequestURL(qValue) {
		jsonErrResponse(c, ErrorInvalidParam, qParams)
		return
	}

	item, err := h.USUsecase.Fetch(c, qValue)
	if err != nil {
		jsonErrResponse(c, err, qParams)
		return
	}

	jsonData(c, http.StatusOK, item)
}

// Insert creates new Short URL
func (h *HTTPURLShortenHandler) Insert(c *gin.Context) {
	var urlSh models.URLShorten
	if c.ShouldBindWith(&urlSh, binding.FormPost) != nil {
		jsonErrResponse(c, ErrorInvalidParam, []string{})
		return
	}

	if !urlsh.IsRequestURL(urlSh.LongURL) {
		return
	}

	res, err := h.USUsecase.Store(c, &urlSh)
	if err != nil {
		jsonErrResponse(c, err, []string{})
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

func jsonErrResponse(c *gin.Context, err error, params []string) {
	status, resp := NewResponseError(c, err, params)
	c.JSON(status, resp)
}
