package openai

import (
	"errors"
)

// ErrorResponse define a error response from the OpenAI API.
type ErrorResponse struct {
	Err struct {
		Message string      `json:"message"`
		Type    string      `json:"type"`
		Param   interface{} `json:"param"`
		Code    interface{} `json:"code"`
	} `json:"error"`
}

// Error returns the error message.
func (r ErrorResponse) Error() string {
	return r.Err.Message
}

const (
	// invalidRequestError is the error type for invalid requests.
	invalidRequestError = "invalid_request_error"
	// invalidRequestError is the error type for invalid requests.
	insufficientQuota = "insufficient_quota"
)

// IsInvalidRequestError returns true if the error is an invalid request error.
func IsInvalidRequestError(err error) bool {
	var respErr ErrorResponse
	if ok := errors.As(err, &respErr); ok {
		return respErr.Err.Type == invalidRequestError
	}
	return false
}

// IsInsufficientQuota returns true if the error is an insufficient quota error.
func IsInsufficientQuota(err error) bool {
	var respErr ErrorResponse
	if ok := errors.As(err, &respErr); ok {
		return respErr.Err.Type == insufficientQuota
	}
	return false
}
