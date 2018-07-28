package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"gopkg.in/mgo.v2/bson"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"urlshortener.api/models"
	urlsh "urlshortener.api/urlshorten"
	"urlshortener.api/urlshorten/db"
	httpDelivery "urlshortener.api/urlshorten/delivery/http"
	"urlshortener.api/urlshorten/usecase"
)

type HTTPTestSuite struct {
	suite.Suite
	memDB  *db.MemoryDB
	router *gin.Engine
	sids   []string
}

// FeedMemDB stores some records in memoryDB
func (suite *HTTPTestSuite) FeedMemDB() {
	// set ctx to confirm Store method
	ctx, _ := gin.CreateTestContext(nil)

	// save generated short IDs for later use
	suite.sids = make([]string, 0)
	sid, _ := shortid.Generate()
	suite.sids = append(suite.sids, sid)

	// data to save in memoryDB
	data := &models.URLShorten{
		ID:        bson.NewObjectId(),
		LongURL:   "https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go",
		ShortURL:  "http://192.168.99.100:8080/" + sid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := suite.memDB.Store(ctx, data)
	if err != nil {
		suite.T().Error(err)
	}
}

func (suite *HTTPTestSuite) SetupSuite() {
	suite.memDB = db.NewMemoryDB()
	usecases := usecase.NewUrlService(suite.memDB)
	handler := httpDelivery.NewUrlHandler(usecases)

	suite.router = SetupRouter(handler)
	suite.FeedMemDB()
}

func (suite *HTTPTestSuite) TearDownSuite() {
	suite.memDB.Close()
	suite.router = nil
	suite.sids = nil
}

func (suite *HTTPTestSuite) TestGet() {
	// load testcases from json file
	path := filepath.Join("urlshorten/test-fixtures", "router_testcases.json")
	tcases := []TestResponse{}

	err := urlsh.ParseFile(path, &tcases)
	if err != nil {
		suite.T().Error(err)
	}

	for _, tc := range tcases {
		var req *http.Request
		if tc.TCID == 3 && len(suite.sids) > 0 { // this TestCase represents the successful response
			req, _ = http.NewRequest("GET", tc.TestCase+suite.sids[0], nil)
		} else {
			req, _ = http.NewRequest("GET", tc.TestCase, nil)
		}

		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)

		if w.Code != 200 {
			var errResp map[string]interface{}
			if err := json.Unmarshal(w.Body.Bytes(), &errResp); err != nil {
				suite.T().Fatal(err)
			}

			assertEqualWithColors(suite, tc.TCID, tc.TestCase, tc.Expected["status"], errResp["status"])
			assertEqualWithColors(suite, tc.TCID, tc.TestCase, tc.Expected["message"], errResp["message"])
		} else {
			var successResp map[string]models.URLShorten
			if err := json.Unmarshal(w.Body.Bytes(), &successResp); err != nil {
				suite.T().Fatal(err)
			}

			if len(suite.sids) > 0 {
				assert.Equal(suite.T(), 200, w.Code)
				assert.Equal(suite.T(), "http://192.168.99.100:8080/"+suite.sids[0], successResp["data"].ShortURL)
				assert.True(suite.T(), successResp["data"].ID.Valid())
			}
		}
	}
}

func (suite *HTTPTestSuite) TestInsert() {
	form := url.Values{}
	form.Add("longUrl", "http://example.com/")

	req, _ := http.NewRequest("POST", "/api/v1/url", strings.NewReader(form.Encode()))
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 201, w.Code)
	assert.Equal(suite.T(), suite.memDB.Size(), 2) // now, memoryDB has 2 records

}

func TestRunHTTPTestSuite(t *testing.T) {
	httpTestsuite := new(HTTPTestSuite)
	suite.Run(t, httpTestsuite)
}

type TestResponse struct {
	TCID     int
	TestCase string
	Expected map[string]interface{}
}

func assertEqualWithColors(s *HTTPTestSuite, tcid int, tc, expected, actual interface{}) {
	assert.Equalf(s.T(), expected, actual, "\n\033[36mTCID\033[0m: %#v \t \033[34mTC\033[0m: \033[35m%#v\033[0m\n\n\033[31m- Expected: %#v \n\033[32m+ Actual: %#v\033[0m \n\n", tcid, tc, expected, actual)
}
