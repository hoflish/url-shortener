package http

import (
	"reflect"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

type (
	RespponseError struct {
		Request map[string]interface{} `json:"request,omitempty"`
		Errors  []*clientError
		Status  int `json:"status,omitempty"`
	}

	clientError struct {
		Code      string      `json:"code,omitempty"`
		Message   string      `json:"message,omitempty"`
		Parameter string      `json:"parameter,omitempty"`
		Value     interface{} `json:"value,omitempty"`
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
	if err == ErrorShortID {
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
func NewResponseError(ctx echo.Context, err error, params []string) error {
	cliErrs := make([]*clientError, 0)
	status, cli := statusCode(err)
	resp := &RespponseError{}
	resp.Status = status

	paramsv := reflect.TypeOf(params)
	if paramsv.Kind() != reflect.Slice {
		panic("params argument must be a slice")
	}

	if status >= 400 && status < 500 {
		if ctx.Request().Method == "POST" || ctx.Request().Method == "PUT" || ctx.Request().Method == "PATCH" {
			formParams, err := ctx.FormParams()
			if err != nil {
				logrus.Error(err)
			}
			resp.Request = map[string]interface{}{
				"form": formParams,
			}
		} else {
			resp.Request = map[string]interface{}{
				"params": ctx.QueryParams(),
			}
		}

		if len(params) > 1 {
			for _, p := range params {
				cli.Message = err.Error()
				cli.Parameter = p
				cli.Value = ctx.QueryParam(p)
				resp.Errors = append(cliErrs, cli)
			}
		} else {
			cli.Message = err.Error()
			cli.Parameter = params[0]
			cli.Value = ctx.QueryParam(params[0])
			resp.Errors = append(cliErrs, cli)
		}
		return ctx.JSON(status, resp)
	}

	cli.Message = err.Error()
	resp.Errors = append(cliErrs, cli)
	return ctx.JSON(status, resp)
}
