package response

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func extractValidationFields(err error) []FieldError {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return nil
	}

	fields := make([]FieldError, 0, len(ve))
	for _, fe := range ve {
		fields = append(fields, FieldError{
			Field:   fe.Field(),
			Message: validationMessage(fe),
		})
	}
	return fields
}

func validationMessage(fe validator.FieldError) string {
	switch fe.Tag() {

	// presence
	case "required":
		return "is required"

	// strings
	case "min":
		return "must be at least " + fe.Param() + " characters"
	case "max":
		return "must be at most " + fe.Param() + " characters"
	case "len":
		return "must be exactly " + fe.Param() + " characters"
	case "oneof":
		return "must be one of: " + fe.Param()
	case "startswith":
		return "must start with " + fe.Param()
	case "endswith":
		return "must end with " + fe.Param()
	case "contains":
		return "must contain " + fe.Param()
	case "excludes":
		return "must not contain " + fe.Param()

	// numbers
	case "gt":
		return "must be greater than " + fe.Param()
	case "gte":
		return "must be greater than or equal to " + fe.Param()
	case "lt":
		return "must be less than " + fe.Param()
	case "lte":
		return "must be less than or equal to " + fe.Param()

	// formats
	case "email":
		return "must be a valid email address"
	case "url":
		return "must be a valid URL"
	case "uuid":
		return "must be a valid UUID"
	case "ip":
		return "must be a valid IP address"
	case "ipv4":
		return "must be a valid IPv4 address"
	case "ipv6":
		return "must be a valid IPv6 address"
	case "hostname":
		return "must be a valid hostname"

	// collections
	case "unique":
		return "must contain unique values"
	case "dive":
		return "contains an invalid value"

	// cross-field
	case "eqfield":
		return "must be equal to " + fe.Param()
	case "nefield":
		return "must not be equal to " + fe.Param()
	case "required_with":
		return "is required with " + fe.Param()
	case "required_without":
		return "is required without " + fe.Param()
	case "required_if":
		return "is required based on condition"

	// time
	case "datetime":
		return "must be a valid datetime (" + fe.Param() + ")"

	// fallback
	default:
		return "is invalid"
	}
}
