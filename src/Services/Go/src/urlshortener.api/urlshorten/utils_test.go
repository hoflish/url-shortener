package urlshorten_test

import (
	"path/filepath"
	"testing"

	"urlshortener.api/urlshorten"
)

type TC struct {
	TCID     int
	TestCase string
	Expected bool
}

func TestIsURL(t *testing.T) {
	path := filepath.Join("test-fixtures", "url_testcases.json")
	tcases := []TC{}

	err := urlshorten.ParseFile(path, &tcases)
	if err != nil {
		t.Error(err)
	}

	for _, tc := range tcases {
		actual := urlshorten.IsRequestURL(tc.TestCase)
		if actual != tc.Expected {
			t.Errorf("\n\033[36mTCID\033[0m: %d \t \033[34mTC\033[0m: \033[35m%v\033[0m\n\n\033[31m- Actual: %t \n\033[32m+ Expected: %t\033[0m \n\n", tc.TCID, tc.TestCase, actual, tc.Expected)
		}
	}
}
