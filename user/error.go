package user

import "errors"

var (
	ErrWrongCredentials = errors.New("wrong credentials")
	ErrEmailInUse       = errors.New("email is already in use")
)
