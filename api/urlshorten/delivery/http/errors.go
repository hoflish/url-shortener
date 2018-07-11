package httphandler

import (
	"errors"
)

// Errors
var (
	ErrorInternalError    = errors.New("Internal Server Error")
	ErrorResourceNotFound = errors.New("Resource Not Found")
	ErrorMissingParam     = errors.New("Missing Paramater")
	ErrorInvalidParam     = errors.New("Validation Failed")
	ErrorUnknown          = errors.New("Unknown Error")
)

// Error Codes
const (
	CodeMissingParam  = "MISSING_PARAMETER"
	CodeInvalidParam  = "INVALID_PARAMETER"
	CodeInternalError = "INTERNAL_ERROR"
	CodeNotFound      = "NOT_FOUND"
	CodeUnknown       = "UNKNOWN"
)
