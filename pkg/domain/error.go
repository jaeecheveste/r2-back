package domain

import "errors"

// ErrNotFound error custom.
var (
	ErrNotFound      = errors.New("not found")
	ErrInternalError = errors.New("internal error")
)
