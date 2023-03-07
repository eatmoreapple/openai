package openai

import (
	"errors"
	"net/http"
	"strings"
)

const (
	// invalidRequestError is the error type for invalid requests.
	invalidRequestError = "invalid_request_error"
	// invalidRequestError is the error type for invalid requests.
	insufficientQuota = "insufficient_quota"
)

// ErrorResponse define a error response from the OpenAI API.
type ErrorResponse struct {
	Err struct {
		Message string      `json:"message"`
		Type    string      `json:"type"`
		Param   interface{} `json:"param"`
		Code    interface{} `json:"code"`
	} `json:"error"`
	StatusCode int
}

// Error returns the error message.
func (r ErrorResponse) Error() string {
	return r.Err.Message
}

// IsInvalidRequest returns true if the error is an invalid request error.
func (r ErrorResponse) IsInvalidRequest() bool {
	return r.Err.Type == invalidRequestError
}

// IsInsufficientQuota returns true if the error is an insufficient quota error.
func (r ErrorResponse) IsInsufficientQuota() bool {
	return r.Err.Type == insufficientQuota
}

// IsRateLimited returns true if the error is a rate limit error.
func (r ErrorResponse) IsRateLimited() bool {
	return r.Err.Type == "requests" && strings.Contains(r.Err.Message, "Rate limit reached")
}

// IsNeedRetryAgain returns true if the error is a need retry again error.
func (r ErrorResponse) IsNeedRetryAgain() bool {
	return r.StatusCode == http.StatusConflict
}

// IsInvalidRequestError returns true if the error is an invalid request error.
func IsInvalidRequestError(err error) bool {
	var respErr ErrorResponse
	if ok := errors.As(err, &respErr); ok {
		return respErr.IsInvalidRequest()
	}
	return false
}

// IsInsufficientQuotaError returns true if the error is an insufficient quota error.
func IsInsufficientQuotaError(err error) bool {
	var respErr ErrorResponse
	if ok := errors.As(err, &respErr); ok {
		return respErr.IsInsufficientQuota()
	}
	return false
}

// IsRateLimitedError returns true if the error is a rate limit error.
func IsRateLimitedError(err error) bool {
	var respErr ErrorResponse
	if ok := errors.As(err, &respErr); ok {
		return respErr.IsRateLimited()
	}
	return false
}

// IsNeedRetryAgainError returns true if the error is a need retry again error.
func IsNeedRetryAgainError(err error) bool {
	var respErr ErrorResponse
	if ok := errors.As(err, &respErr); ok {
		return respErr.IsNeedRetryAgain()
	}
	return false
}
