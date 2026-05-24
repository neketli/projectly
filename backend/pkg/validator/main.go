package validator

import (
	"projectly-server/pkg/apierror"

	"github.com/go-playground/validator/v10"
)

// Validator provides request validation.
type Validator struct {
	validator *validator.Validate
}

// New creates a new Validator.
func New() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validate validates a struct.
func (cv *Validator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return apierror.Validation(err.Error())
	}
	return nil
}
