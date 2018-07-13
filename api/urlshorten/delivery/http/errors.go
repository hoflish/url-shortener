package httphandler

import (
	"errors"

	"github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

type APIError struct {
	Err     map[string]interface{} `json:"error"`
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

// Errors
var (
	ErrNotFound = errors.New("Resource not found")
	ErrInternal = errors.New("Oops! something went wrong")
	ErrUnknown  = errors.New("An unexpected error occurred")

	ErrAPINotFound = NewAPIError(404, CodeNotFound, ErrNotFound)
	ErrAPIInternal = NewAPIError(500, CodeInternalError, ErrInternal)
	ErrAPIUnknown  = NewAPIError(520, CodeUnknown, ErrUnknown)
)

// Error Codes
const (
	CodeInvalidParam  = "invalidParameter"
	CodeInternalError = "internalError"
	CodeNotFound      = "notFound"
	CodeUnknown       = "unknownError"
	CodeMissingParam  = "missingParameter"
)

func NewAPIError(status int, code string, err error) *APIError {
	errs := []*Error{}
	e := &Error{
		Code:    code,
		Message: err.Error(),
		Details: "",
	}
	errs = append(errs, e)
	return &APIError{
		Err: map[string]interface{}{
			"errors": errs,
		},
		Status:  status,
		Message: e.Message,
	}
}

// IsDBError checks if "err" arg is a MongoDB error or not
func IsDBError(err error) bool {
	switch e := err.(type) {
	case *mgo.QueryError:
		logrus.Error("DB: QueryError: %v, code: %d", e.Message, e.Code)
		return true
	case *mgo.LastError:
		logrus.Error("DB: LastError: %v, code: %d", e.Err, e.Code)
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
