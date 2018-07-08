package http

import (
	"errors"
)

// Errors
var (
	ErrorInternalError    = errors.New("Internal Server Error")
	ErrorShortID          = errors.New("Could not generate short ID")
	ErrorResourceNotFound = errors.New("Resource Not Found")
	ErrorMissingParam     = errors.New("Missing Paramater")
	ErrorInvalidParam     = errors.New("Validation Failed")
	ErrorUnknown          = errors.New("Unknown Error")
)

// Error Codes
const (
	CodeSuccess       = "SUCCESS"
	CodeMissingParam  = "MISSING_PARAMETER"
	CodeInvalidParam  = "INVALID_PARAMETER"
	CodeInternalError = "INTERNAL_ERROR"
	CodeReqBodyEmpty  = "REQUEST_BODY_EMPTY"
	CodeNotFound      = "NOT_FOUND"
	Unknown           = "Unknown"
)
