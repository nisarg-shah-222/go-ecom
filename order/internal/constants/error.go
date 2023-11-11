package constants

import "errors"

var (
	ErrBadRequest          error = errors.New("bad request")
	ErrInternalServerError error = errors.New("internal server error")

	ErrValidationFailed error = errors.New("validation failed")
	ErrOrderNotExists   error = errors.New("order does not exists")
	ErrProductAPIFailed error = errors.New("product api failed")
)
