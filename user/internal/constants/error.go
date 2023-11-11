package constants

import "errors"

var (
	ErrBadRequest          error = errors.New("bad request")
	ErrInternalServerError error = errors.New("internal server error")

	ErrUserNameExists    error = errors.New("username already exists")
	ErrUserNameNotExists error = errors.New("username does not exists")
	ErrIncorrectPassword error = errors.New("incorrect password")
)
