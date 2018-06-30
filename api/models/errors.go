package models

import "errors"

var (
	ErrorNotFound          = errors.New("URLShorten not found")
	ErrorMissingQueryParam = errors.New("Missing shortUrl query parameter")
	ErrorInvalidURL        = errors.New("Invalid shortUrl")
	ErrorInternalServer    = errors.New("Internal Server Error")
)
