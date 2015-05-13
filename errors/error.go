package errors

import (
	"fmt"
)

const (
	// Agent
	ErrInvalidUserInput ErrorCode = "Agent: Invalid user inputs."
	ErrInvalidGoal      ErrorCode = "Agent: Invalid goal."
	ErrActionNotFound   ErrorCode = "Agent: No action found for the given goal."

	// Arg parsing
	ErrInvalidInput  ErrorCode = "Args: Invalid input."
	ErrInvalidParams ErrorCode = "Args: Invalid params. No params retreived from user input."
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
