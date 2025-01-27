package serviceerror

type ErrorMessage string

var (
	// General
	ServerError    ErrorMessage = "errors.serverError"
	RecordNotFound ErrorMessage = "errors.recordNotFound"
	NoRowsEffected ErrorMessage = "errors.noRowsEffected"

	// Validation
	InvalidRequestBody ErrorMessage = "errors.invalidRequestBody"
)
