package factory

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
	MakeFromName(name string) *Command
	MakeFromMetadata(meta Metadata) *Command
}

func NewFactory() (factory *FactoryImp) {
	return new(FactoryImp)
}

func MakeFromName(name string) Command {
	return NewCommand(name)
}
