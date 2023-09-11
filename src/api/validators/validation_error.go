package validators

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}

func GetValidationError(err error) *[]ValidationError {
	var validationErrors []ValidationError
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, e := range err.(validator.ValidationErrors) {
			var el ValidationError
			el.Property = e.Field()
			el.Tag = e.Tag()
			el.Value = e.Param()
			validationErrors = append(validationErrors, el)
		}
		return &validationErrors
	}
	return nil
}
