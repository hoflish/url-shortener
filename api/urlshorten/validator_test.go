package urlshorten_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/hoflish/url-shortener/api/urlshorten"
)

type TC struct {
	TCID     int
	TestCase string
	Expected bool
}

func parseFile(file string) ([]TC, error) {
	tc := make([]TC, 0)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &tc)
	return tc, err
}

func TestIsURL(t *testing.T) {
	path := filepath.Join("test-fixtures", "url_testcases.json")
	cases, _ := parseFile(path)

	for _, tc := range cases {
		actual := urlshorten.IsRequestURL(tc.TestCase)
		if actual != tc.Expected {
			t.Errorf("\n\033[36mTCID\033[0m: %d \t \033[34mTC\033[0m: \033[35m%v\033[0m\n\n\033[31m- Actual: %t \n\033[32m+ Expected: %t\033[0m \n\n", tc.TCID, tc.TestCase, actual, tc.Expected)
		}
	}
}
