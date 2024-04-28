package errors

import "fmt"

type Error struct {
	Code    string
	Message string
}

func NewError(code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("Code: %s. Message: %s", e.Code, e.Message)
}

var IncorrectReqBodyErr = Error{
	Code:    "INCORRECT_REQUEST_BODY",
	Message: "Failed to parse request body",
}

var ValidationFailedErr = Error{
	Code:    "VALIDATION_FAILED",
	Message: "Failed to validate parameters",
}

var InvalidParamEterErr = Error{
	Code:    "INVALID_PARAMETER",
	Message: "Invalid parameter",
}
