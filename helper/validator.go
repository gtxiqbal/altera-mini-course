package helper

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
	errorMessage  = map[string]string{
		"email":    "email invalid",
		"required": "required",
	}
)

type FieldErrorResponse map[string]string
type FieldErrorResponses []FieldErrorResponse

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (valid *CustomValidator) Validate(i interface{}) error {
	if err := valid.Validator.Struct(i); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			fieldErrorResponses := make(FieldErrorResponses, len(validationErrors))
			for i, ve := range validationErrors {
				field := toSnakeCase(ve.Field())
				fieldErrorResponses[i] = FieldErrorResponse{
					field: errorMessage[ve.Tag()],
				}
			}
			return &fieldErrorResponses
		}
	}
	return nil
}

func (f *FieldErrorResponses) Error() string {
	return fmt.Sprintf("invalid request")
}
