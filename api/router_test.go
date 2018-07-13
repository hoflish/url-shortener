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

	"github.com/hoflish/url-shortener/api/models"
	urlsh "github.com/hoflish/url-shortener/api/urlshorten"
	"github.com/hoflish/url-shortener/api/urlshorten/db"
	"github.com/hoflish/url-shortener/api/urlshorten/delivery/http"
	"github.com/hoflish/url-shortener/api/urlshorten/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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
	usecases := usecase.NewURLShortenUsecase(suite.memDB)
	handler := httphandler.NewHTTPURLShortenHandler(usecases)

	suite.router = setupRouter(handler)
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
	cases := []TestResponse{}
	err := urlsh.ParseFile(path, &cases)

	if err != nil {
		suite.T().Error(err)
	}

	for _, tc := range cases {
		var req *http.Request
		w := httptest.NewRecorder()

		if tc.TCID == 3 && len(suite.sids) > 0 { // this TestCase represents the successful response
			req, _ = http.NewRequest("GET", tc.TestCase+suite.sids[0], nil)
		} else {
			req, _ = http.NewRequest("GET", tc.TestCase, nil)
		}
		suite.router.ServeHTTP(w, req)

		if w.Code != 200 {
			var errResp map[string]interface{}
			if err := json.Unmarshal(w.Body.Bytes(), &errResp); err != nil {
				suite.T().Fatal(err)
			}
			assert.Equal(suite.T(), tc.Expected["status"], errResp["status"])
			assert.Equal(suite.T(), tc.Expected["message"], errResp["message"])
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

	req, _ := http.NewRequest("POST", "/api/url", strings.NewReader(form.Encode()))
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
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
