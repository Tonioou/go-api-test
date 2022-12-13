package model

import (
	"errors"
	"net/http"
	"time"

	"github.com/joomcode/errorx"
)

var (
	namespace = errorx.NewNamespace("common")
	NotFound  = errorx.NewType(namespace, "not_found", errorx.RegisterTrait("not_found"))
)

type ErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`

	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type logFunc func(...interface{})

func NewErrorResponse(err error, log logFunc) *ErrorResponse {
	log(err.Error())
	errorResponse := &ErrorResponse{}
	errorResponse.fillProperties(err)
	return errorResponse
}

func (er *ErrorResponse) fillProperties(err error) {
	statusCode := 0
	var errx *errorx.Error
	if isErrorx := errors.As(err, &errx); isErrorx {
		switch {
		case errx.IsOfType(NotFound):
			statusCode = http.StatusNotFound
		case errx.IsOfType(errorx.InternalError):
			statusCode = http.StatusInternalServerError
		case errx.IsOfType(errorx.IllegalArgument):
			statusCode = http.StatusBadRequest
		case errx.IsOfType(errorx.NotImplemented):
			statusCode = http.StatusNotImplemented
		default:
			statusCode = http.StatusInternalServerError
		}

		er.StatusCode = statusCode

	} else {
		er.StatusCode = http.StatusInternalServerError
	}
	er.Message = errx.Error()
	er.Timestamp = time.Now()

}
