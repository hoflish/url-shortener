package httphandler

import (
	"net/url"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

type (
	// ResponseError represents the typical response structure
	// for a given error
	ResponseError struct {
		Request *request       `json:"request"`
		Errors  []*clientError `json:"errors"`
		Status  int            `json:"status"`
	}

	request struct {
		Params        url.Values `json:"params"`
		OperationType string     `json:"operation_type"`
	}

	clientError struct {
		Code      string `json:"code"`
		Message   string `json:"message"`
		Parameter string `json:"parameter"`
		Value     string `json:"value"`
	}
)

func isInternalError(err error) bool {
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

func isNotFound(err error) bool {
	if err == mgo.ErrNotFound || err == ErrorResourceNotFound {
		return true
	}
	return false
}

func isValidationError(err error) bool {
	if err == ErrorInvalidParam {
		return true
	}
	return false
}

func isBadRequestError(err error) bool {
	if err == ErrorMissingParam {
		return true
	}
	return false
}

func statusCode(err error) (int, *clientError) {
	status := 0
	cli := &clientError{}

	if isInternalError(err) {
		cli.Code = CodeInternalError
		status = 500
	}
	if isNotFound(err) {
		cli.Code = CodeNotFound
		status = 404
	}
	if isBadRequestError(err) {
		cli.Code = CodeMissingParam
		status = 400
	}
	if isValidationError(err) {
		cli.Code = CodeInvalidParam
		status = 422
	}
	return status, cli
}

// NewResponseError returns a formatted error Response
func NewResponseError(ctx *gin.Context, err error, params []string) (int, *ResponseError) {
	paramsv := reflect.TypeOf(params)
	if paramsv.Kind() != reflect.Slice {
		panic("params argument must be a slice")
	}

	cliErrs := make([]*clientError, 0)
	status, cli := statusCode(err)
	resp := &ResponseError{}
	resp.Status = status

	if status >= 400 && status < 500 {
		if ctx.Request.Method == "GET" {

		} else {
			req := &request{Params: ctx.Request.URL.Query()}
			resp.Request = req
		}

		if len(params) > 1 {
			for _, p := range params {
				cli.Message = err.Error()
				cli.Parameter = p
				cli.Value = ctx.Query(p)
				resp.Errors = append(cliErrs, cli)
			}
		} else {
			cli.Message = err.Error()
			cli.Parameter = params[0]
			cli.Value = ctx.Query(params[0])
			resp.Errors = append(cliErrs, cli)
		}
		return status, resp
	}

	cli.Message = err.Error()
	resp.Errors = append(cliErrs, cli)
	return status, resp
}
