package urlshorten

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
)

// ParseFile parses the JSON file and stores JSON-encoded data
// in the value pointed to by v
func ParseFile(file string, v interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &v)
}

// IsRequestURL check if the string rawurl, assuming
// it was received in an HTTP request, is a valid
// URL confirm to RFC 3986
// FIXME: Use a consistent URL validator as the Client
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
