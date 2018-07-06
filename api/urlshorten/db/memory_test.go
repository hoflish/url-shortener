package db_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/teris-io/shortid"
	"gopkg.in/mgo.v2/bson"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/hoflish/url-shortener/api/models"
	"github.com/hoflish/url-shortener/api/urlshorten"

	"github.com/hoflish/url-shortener/api/urlshorten/db"
)

type DBTestSuite struct {
	suite.Suite
	MemDB *db.MemoryDB
	SID   string
}

// The SetupSuite method will be run by testify once, at the very
// start of the testing suite, before any tests are run.
func (suite *DBTestSuite) SetupSuite() {
	suite.MemDB = db.NewMemoryDB()
}

// The TearDownSuite method will be run by testify once, at the very
// end of the testing suite, after all tests have been run.
func (suite *DBTestSuite) TearDownSuite() {
	suite.MemDB.Close()
	suite.SID = ""
}

// Every method in a testing suite that begins with "Test" will be run
// as a test.
func (suite *DBTestSuite) TestAStore() {
	sid, err := shortid.Generate()
	suite.SID = sid
	// Assert that no error returned
	if assert.NoError(suite.T(), err) {
		// Assert that sid length is 9 and of type string
		assert.Len(suite.T(), sid, 9)
		assert.IsType(suite.T(), "", sid)
	}

	// Create URLShorten resource
	data := &models.URLShorten{
		ID:        bson.NewObjectId(),
		LongURL:   "https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go",
		ShortURL:  "http://192.168.99.100:8080/" + sid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Target Store method that saves data into in-memory DB
	res, _ := suite.MemDB.Store(context.TODO(), data)

	// Assert that ShortURL is a Valid URL
	assert.True(suite.T(), urlshorten.IsRequestURL(res.ShortURL), "wanted \033[31m'%s'\033[0m to be Valid URL", res.ShortURL)

	// Assert that repos size is 1 after the first store
	assert.Equal(suite.T(), suite.MemDB.Size(), 1)

	// Assert ID length is 12 - genrated by bson.NewObjectId()
	assert.Len(suite.T(), res.ID, 12)
}

func (suite *DBTestSuite) TestBFetch() {
	shortURLs := []string{
		"http://192.168.99.100:8080/foobarzoo",
		"http://192.168.99.100:8080/" + suite.SID,
	}

	for _, v := range shortURLs {
		res, err := suite.MemDB.Fetch(context.TODO(), v)

		if err != nil {
			expectedError := fmt.Sprintf("Memorydb: URLShorten not found with shortURL %s", v)
			assert.Equal(suite.T(), expectedError, err.Error())
		}

		if res != nil {
			assert.Equal(suite.T(), res.ShortURL[27:], suite.SID)
		}
	}

	// Empty the repos
	suite.MemDB.Close()

	// Assert that repos size is 0
	assert.Equal(suite.T(), suite.MemDB.Size(), 0)

}

func TestRunDBTestSuite(t *testing.T) {
	dbtestsuite := new(DBTestSuite)
	suite.Run(t, dbtestsuite)

	assert.Equal(t, dbtestsuite.SID, "")
}
