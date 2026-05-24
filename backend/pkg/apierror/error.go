package apierror

import (
	"errors"
	"net/http"
)

// Kind categorizes the type of an API error.
type Kind int

const (
	// KindNotFound indicates the requested resource was not found.
	KindNotFound Kind = iota + 1
	// KindValidation indicates invalid input.
	KindValidation
	// KindUnauthorized indicates missing or invalid authentication.
	KindUnauthorized
	// KindForbidden indicates the user lacks permission.
	KindForbidden
	// KindConflict indicates a resource conflict (e.g., duplicate).
	KindConflict
	// KindInternal indicates an unexpected server error.
	KindInternal
)

var kindStrings = map[Kind]string{
	KindNotFound:     "NOT_FOUND",
	KindValidation:   "VALIDATION_ERROR",
	KindUnauthorized: "UNAUTHORIZED",
	KindForbidden:    "FORBIDDEN",
	KindConflict:     "CONFLICT",
	KindInternal:     "INTERNAL_ERROR",
}

func (k Kind) String() string {
	if s, ok := kindStrings[k]; ok {
		return s
	}
	return "UNKNOWN"
}

// HTTPStatus returns the HTTP status code for the error kind.
func (k Kind) HTTPStatus() int {
	switch k {
	case KindNotFound:
		return http.StatusNotFound
	case KindValidation:
		return http.StatusBadRequest
	case KindUnauthorized:
		return http.StatusUnauthorized
	case KindForbidden:
		return http.StatusForbidden
	case KindConflict:
		return http.StatusConflict
	case KindInternal:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

// KindFromHTTPStatus maps an HTTP status code to the closest error Kind.
func KindFromHTTPStatus(code int) Kind {
	switch code {
	case http.StatusBadRequest:
		return KindValidation
	case http.StatusUnauthorized:
		return KindUnauthorized
	case http.StatusForbidden:
		return KindForbidden
	case http.StatusNotFound:
		return KindNotFound
	case http.StatusConflict:
		return KindConflict
	case http.StatusInternalServerError:
		return KindInternal
	default:
		return KindInternal
	}
}

// Error is a domain error with a Kind, a human-readable message, and an optional wrapped error.
type Error struct {
	Kind    Kind
	Message string
	Err     error
}

func (e *Error) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func (e *Error) Unwrap() error {
	return e.Err
}

// KindOf extracts the Kind from an error, defaulting to KindInternal.
func KindOf(err error) Kind {
	var e *Error
	if errors.As(err, &e) {
		return e.Kind
	}
	return KindInternal
}

// New creates a new categorized error with the given kind and message.
func New(kind Kind, message string) *Error {
	return &Error{Kind: kind, Message: message}
}

// Wrap creates a new categorized error wrapping an internal error.
func Wrap(kind Kind, message string, err error) *Error {
	return &Error{Kind: kind, Message: message, Err: err}
}

// NotFound creates a 404 error.
func NotFound(message string) *Error {
	return New(KindNotFound, message)
}

// Validation creates a 400 error.
func Validation(message string) *Error {
	return New(KindValidation, message)
}

// Unauthorized creates a 401 error.
func Unauthorized(message string) *Error {
	return New(KindUnauthorized, message)
}

// Forbidden creates a 403 error.
func Forbidden(message string) *Error {
	return New(KindForbidden, message)
}

// Conflict creates a 409 error.
func Conflict(message string) *Error {
	return New(KindConflict, message)
}

// Internal creates a 500 error.
func Internal(message string) *Error {
	return New(KindInternal, message)
}

// NotFoundWrap creates a 404 error wrapping an internal cause.
func NotFoundWrap(message string, err error) *Error {
	return Wrap(KindNotFound, message, err)
}

// ValidationWrap creates a 400 error wrapping an internal cause.
func ValidationWrap(message string, err error) *Error {
	return Wrap(KindValidation, message, err)
}

// UnauthorizedWrap creates a 401 error wrapping an internal cause.
func UnauthorizedWrap(message string, err error) *Error {
	return Wrap(KindUnauthorized, message, err)
}

// ForbiddenWrap creates a 403 error wrapping an internal cause.
func ForbiddenWrap(message string, err error) *Error {
	return Wrap(KindForbidden, message, err)
}

// ConflictWrap creates a 409 error wrapping an internal cause.
func ConflictWrap(message string, err error) *Error {
	return Wrap(KindConflict, message, err)
}

// InternalWrap creates a 500 error wrapping an internal cause.
func InternalWrap(message string, err error) *Error {
	return Wrap(KindInternal, message, err)
}
