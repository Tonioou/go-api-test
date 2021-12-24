package model

import (
	"net/http"
	"time"

	"github.com/joomcode/errorx"
)

var (
	namespace = errorx.NewNamespace("common")
	NotFound  = errorx.NewType(namespace, "not_found", errorx.RegisterTrait("not_found"))
)

type ErrorResponse struct {
	StatusCode int       `json:"status_code"`
	Message    string    `json:"message"`
	Timestamp  time.Time `json:"timestamp"`
}

type logFunc func(...interface{})

func NewErrorResponse(errx *errorx.Error, log logFunc) *ErrorResponse {
	log(errx.Error())
	errorResponse := &ErrorResponse{}
	errorResponse.fillProperties(errx)
	return errorResponse
}

func (er *ErrorResponse) fillProperties(errx *errorx.Error) {
	statusCode := 0
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
	er.Message = errx.Error()
	er.Timestamp = time.Now()
}
