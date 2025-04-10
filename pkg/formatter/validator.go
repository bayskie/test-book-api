package formatter

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func NewValidationErrors(err error) *string {
	var errors []string

	if err == nil {
		return nil
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			var msg string

			field := strings.ToLower(fieldErr.Field())
			tag := strings.ToLower(fieldErr.Tag())

			switch tag {
			case "required":
				msg = fmt.Sprintf("%s is required", field)
			case "min":
				msg = fmt.Sprintf("%s must be at least %s characters", field, strings.ToLower(fieldErr.Param()))
			case "max":
				msg = fmt.Sprintf("%s must be at most %s characters", field, strings.ToLower(fieldErr.Param()))
			case "gte":
				msg = fmt.Sprintf("%s must be greater than or equal to %s", field, strings.ToLower(fieldErr.Param()))
			case "lte":
				msg = fmt.Sprintf("%s must be less than or equal to %s", field, strings.ToLower(fieldErr.Param()))
			default:
				msg = fmt.Sprintf("%s is not valid", field)
			}

			errors = append(errors, msg)
		}
	} else {
		errors = append(errors, err.Error())
	}

	ret := strings.Join(errors, ", ")

	return &ret
}
