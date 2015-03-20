package command

import (
	"fmt"
)

var (
	errInvalidNameCode ErrorCode = "[Factory] Invalid command name"
)

type FactoryError struct {
	code ErrorCode
	msg  string
}

func (e *FactoryError) Error() string {
	return fmt.Sprintf("%s : expr=<%v>", e.code, e.expr)
}

type Factory interface {
	MakeFromName(name string) *CommandImp
	MakeFromMetadata(meta Metadata) *CommandImp
}

func New() (factory *FactoryImp) {
	return new(FactoryImp)
}
