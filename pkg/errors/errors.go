package errors

import "errors"

var (
	ShortPassword    = errors.New("password must contain at least 8 characters")
	UserAlreadyExist = errors.New("user with this login already exists")
	NullForeignKey   = errors.New("cannot bind foreignKey to null ID")
)
