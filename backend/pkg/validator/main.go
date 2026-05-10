package validator

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
