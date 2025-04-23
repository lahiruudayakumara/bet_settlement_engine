package errors

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrEmailExists       = errors.New("email already registered")
	ErrInvalidEmail      = errors.New("invalid email format")
	ErrNegativeBalance   = errors.New("balance cannot be negative")
	ErrUserAlreadyExists = errors.New("user already exists")
)
