package errors

import "errors"

var (
	ShortPassword = errors.New("password must contain at least 8 characters")
)
