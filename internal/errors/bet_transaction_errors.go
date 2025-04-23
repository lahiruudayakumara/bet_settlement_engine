package errors

import "errors"

var (
	ErrTransactionNotFound     = errors.New("transaction not found")
	ErrInvalidTransactionType  = errors.New("invalid transaction type")
	ErrDuplicateTransactionID  = errors.New("transaction ID already exists")
	ErrTransactionFailed       = errors.New("failed to process transaction")
)