package handler

import (
	"net/http"
	"short-link/pkg/serviceerror"
)

// StatusCodeMapping maps error to http status code
var StatusCodeMapping = map[serviceerror.ErrorMessage]int{
	// General
	serviceerror.ServerError:    http.StatusInternalServerError,
	serviceerror.RecordNotFound: http.StatusNotFound,
	serviceerror.NoRowsEffected: http.StatusNotFound,
	// Validation
	serviceerror.InvalidRequestBody: http.StatusBadRequest,
}
