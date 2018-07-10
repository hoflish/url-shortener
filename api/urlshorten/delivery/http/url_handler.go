package httphandler

import (
	"net/http"

	models "github.com/hoflish/url-shortener/api/models"
	urlsh "github.com/hoflish/url-shortener/api/urlshorten"
	"github.com/labstack/echo"
)

type HTTPURLShortenHandler struct {
	USUsecase urlsh.URLShortenUsecase
}

// Get method gets information for a specified short URL
func (h *HTTPURLShortenHandler) Get(c echo.Context) error {
	qparam := "shortUrl" // required query param
	query := c.QueryParams()
	qparams := []string{qparam}

	if _, ok := query[qparam]; !ok {
		return NewResponseError(c, ErrorMissingParam, qparams)
	}

	shortURL := query.Get(qparam)
	if !urlsh.IsRequestURL(shortURL) {
		return NewResponseError(c, ErrorInvalidParam, qparams)
	}

	ctx := c.Request().Context()
	item, err := h.USUsecase.Fetch(ctx, shortURL)
	if err != nil {
		return NewResponseError(c, err, qparams)
	}

	return jsonData(c, http.StatusOK, item)
}

// Insert creates new Short URL
func (h *HTTPURLShortenHandler) Insert(c echo.Context) error {
	var urlShorten models.URLShorten
	formParam := "longUrl" // required form value
	formParams := []string{formParam}

	err := c.Bind(&urlShorten)
	if err != nil {
		return NewResponseError(c, err, formParams)
	}

	if !urlsh.IsRequestURL(urlShorten.LongURL) {
		return NewResponseError(c, ErrorInvalidParam, formParams)
	}

	ctx := c.Request().Context()
	res, err := h.USUsecase.Store(ctx, &urlShorten)
	if err != nil {
		return NewResponseError(c, err, formParams)
	}

	return jsonData(c, http.StatusOK, res)
}

// NewHTTPURLShortenHandler defines API endpoints
func NewHTTPURLShortenHandler(e *echo.Echo, u urlsh.URLShortenUsecase) *HTTPURLShortenHandler {
	return &HTTPURLShortenHandler{
		USUsecase: u,
	}
}

func jsonData(c echo.Context, status int, data interface{}) error {
	d := map[string]interface{}{
		"data": data,
	}
	return c.JSON(status, d)
}
