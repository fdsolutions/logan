package errors

import (
	"fmt"
)

const (
	// Agent
	ErrInvalidUserInput ErrorCode = "Invalid user inputs."

	// Arg parsing
	ErrInvalidInput  ErrorCode = "Invalid input."
	ErrInvalidParams ErrorCode = "Invalid params : No params retreived from user input."
)

// ErrorCode is a generic error type
type ErrorCode string

type LoganError interface {
	error
	Code() ErrorCode
}

type loganErrorImpl struct {
	code ErrorCode
}

// NewError create parser error for a specific code
func New(code ErrorCode) LoganError {
	e := &loganErrorImpl{code}
	return LoganError(e)
}

func (err *loganErrorImpl) Error() string {
	return fmt.Sprintf("%s", err.Code)
}

func (err *loganErrorImpl) Code() ErrorCode {
	return err.code
}
