package url

import (
	"net/url"
)

// IsRequestURL check if the string rawurl, assuming
// it was received in an HTTP request, is a valid
// URL confirm to RFC 3986
func IsRequestURL(rawurl string) bool {
	url, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return false
	}
	if len(url.Scheme) == 0 {
		return false
	}
	return true
}
