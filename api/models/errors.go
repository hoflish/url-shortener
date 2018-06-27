package models

import "errors"

var (
	INTERNAL_SERVER_ERROR = errors.New("Internal Server Error")
	NOT_FOUND_ERROR       = errors.New("Url not found")
	MISSING_QUERY_PARAM   = errors.New("Missing shortUrl query parameter")
	INVALID_URL           = errors.New("Invalid shortUrl")
)
