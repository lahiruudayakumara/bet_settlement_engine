package errors

import "fmt"

// BetError defines custom errors related to betting.
type BetError struct {
	Code    int
	Message string
}

func (e *BetError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

const (
	ErrBetUserNotFound    = 1
	ErrInsufficientFunds  = 2
	ErrBetPlacementFailed = 3
)

// NewBetError creates a new BetError.
func NewBetError(code int, message string) *BetError {
	return &BetError{Code: code, Message: message}
}
