package repository

import (
	"gopkg.in/mgo.v2"
)

var (
	IsDrop = true
)

// Init creates the main session to our mongodb instance
func Init(host string) (*mgo.Session, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	// Drop Database
	if IsDrop {
		// TODO: single source of truth
		err = session.DB("url-shortener").DropDatabase()
		if err != nil {
			return nil, err
		}
	}

	// Collection Url
	c := session.DB("url-shortener").C("urlshorten")

	// Index
	index := mgo.Index{
		Key:        []string{"urlshorten"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		return nil, err
	}

	return session, nil
}
