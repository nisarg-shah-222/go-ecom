package constants

import "errors"

var (
	ErrBadRequest          error = errors.New("bad request")
	ErrInternalServerError error = errors.New("internal server error")

	ErrProductExists    error = errors.New("product already exists")
	ErrProductNotExists error = errors.New("product does not exists")

	ErrProductDetailNotExists error = errors.New("product detail not exists")

	ErrPermissionExists error = errors.New("permission already exists")
)
