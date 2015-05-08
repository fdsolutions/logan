package agent

import (
	"github.com/fdsolutions/logan/errors"
)

type StatusCode int

const (
	StatusSuccess StatusCode = iota // 0
	StatusFail
)

type Status interface {
	GetCode() StatusCode
	GetValue() interface{}
	StackError(errors.LoganError)
	GetErrorStackCodes() []errors.ErrorCode
}

type statusImpl struct {
	code        StatusCode
	value       interface{}
	errorsStack []errors.LoganError
}

func (s *statusImpl) GetCode() StatusCode {
	return s.code
}

func (s *statusImpl) GetValue() interface{} {
	return s.value
}

func (s *statusImpl) StackError(err errors.LoganError) {
	s.errorsStack = append(s.errorsStack, err)
}

func (s *statusImpl) GetErrorStackCodes() []errors.ErrorCode {
	codes := make([]errors.ErrorCode, len(s.errorsStack))
	for _, err := range s.errorsStack {
		if loganErr, ok := err.(errors.LoganError); ok {
			codes = append(codes, loganErr.Code())
		}
	}
	return codes
}

func NewStatus(code StatusCode, value interface{}) (s Status) {
	errStack := []errors.LoganError{}
	s = &statusImpl{code, value, errStack}
	return
}
