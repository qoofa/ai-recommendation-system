package response

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	appErr "github.com/qoofa/AI-Recommendation-System/internal/core/errors"
)

type APIError struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Fields  []FieldError `json:"fields,omitempty"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func Error(w http.ResponseWriter, err error) {
	var ae *appErr.AppError

	if e, ok := err.(*appErr.AppError); ok {
		fmt.Fprint(os.Stderr, e.Err)
	}

	if errors.As(err, &ae) {
		writeAppError(w, ae)
		return
	}

	JSON(w, http.StatusInternalServerError, APIResponse{
		Success: false,
		Error: &APIError{
			Code:    string(appErr.Internal),
			Message: "internal server error",
		},
	})
}

func writeAppError(w http.ResponseWriter, err *appErr.AppError) {
	status := mapKindToStatus(err.Kind)

	apiErr := &APIError{
		Code:    string(err.Kind),
		Message: err.Message,
	}

	if fields := extractValidationFields(err.Err); len(fields) > 0 {
		apiErr.Fields = fields
	}

	JSON(w, status, APIResponse{
		Success: false,
		Error:   apiErr,
	})
}

func mapKindToStatus(kind appErr.Kind) int {
	switch kind {
	case appErr.BadRequest:
		return http.StatusBadRequest
	case appErr.NotFound:
		return http.StatusNotFound
	case appErr.Unauthorized:
		return http.StatusUnauthorized
	case appErr.Forbidden:
		return http.StatusForbidden
	case appErr.Conflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
