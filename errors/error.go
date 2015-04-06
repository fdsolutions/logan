package errors

import (
	"fmt"
)

// ErrorCode is a generic error type
type ErrorCode string

type LoganError struct {
	code ErrorCode
}

// NewError create parser error for a specific code
func New(code ErrorCode) *LoganError {
	return &LoganError{code}
}

func (err *LoganError) Error() string {
	return fmt.Sprintf("%s", err.code)
}
