package task

import "errors"

var (
	ErrInvalidPriority = errors.New("invalid priority")
	ErrInvalidToken    = errors.New("invalid token")
)
