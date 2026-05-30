package errors

import (
	stderrors "errors"
	"net/http"
)

type HTTPError struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func ToHTTP(err error) (int, HTTPError) {
	if err == nil {
		return http.StatusOK, HTTPError{}
	}

	var appErr *Error
	if !stderrors.As(err, &appErr) {
		return http.StatusInternalServerError, HTTPError{
			Code:    CodeInternal,
			Message: "internal server error",
		}
	}

	return statusFor(appErr.Code), HTTPError{
		Code:    appErr.Code,
		Message: appErr.Message,
	}
}

func statusFor(code Code) int {
	switch code {
	case CodeInvalidInput:
		return http.StatusBadRequest
	case CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeForbidden:
		return http.StatusForbidden
	case CodeNotFound:
		return http.StatusNotFound
	case CodeConflict, CodeDuplicateTransaction:
		return http.StatusConflict
	case CodeInsufficientPoints, CodeRewardExpired, CodeConsentRequired:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}
