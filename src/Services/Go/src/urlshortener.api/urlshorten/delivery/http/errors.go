package http

import (
	"errors"

	"github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

type APIError struct {
	Err     *Error `json:"error,omitempty"`
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type Error struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Details string `json:"details,omitempty"`
}

// Errors
var (
	ErrNotFound = errors.New("Resource not found")
	ErrInternal = errors.New("Oops! something went wrong")
	ErrUnknown  = errors.New("An unexpected error occurred")

	ErrAPINotFound = NewAPIError(404, CodeNotFound, "", ErrNotFound)
	ErrAPIInternal = NewAPIError(500, CodeInternalError, "", ErrInternal)
)

// Error Codes
const (
	CodeInvalidParam  = "invalidParameter"
	CodeInternalError = "internalError"
	CodeNotFound      = "notFound"
	CodeMissingParam  = "missingParameter"
	CodeBadRequest    = "badRequest"
)

func NewAPIError(status int, code, detail string, err error) *APIError {
	e := &Error{
		Code:    code,
		Message: err.Error(),
		Details: detail,
	}

	return &APIError{
		Err:     e,
		Status:  status,
		Message: e.Message,
	}
}

// IsDBError checks if "err" arg is a MongoDB error or not
func IsDBError(err error) bool {
	switch e := err.(type) {
	case *mgo.QueryError:
		logrus.WithFields(logrus.Fields{
			"msg":  e.Message,
			"code": e.Code,
		}).Error("MongoDB: QueryError")
		return true
	case *mgo.LastError:
		logrus.WithFields(logrus.Fields{
			"msg":  e.Err,
			"code": e.Code,
		}).Error("MongoDB: LastError")
		return true
	}
	return false
}

// IsNotFound checks if "err" arg is a "Not Found" error
func IsNotFound(err error) bool {
	if err == mgo.ErrNotFound || err == ErrNotFound {
		return true
	}
	return false
}
