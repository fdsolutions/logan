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

type Repository interface {
	findCmdByName(name string) (Metadata, error)
}

type Repositories []Repository

type Factory interface {
	getCommandByName(name string) Command
	getCommandByMetadata(meta Metadata) Command
}

type FactoryImp struct {
	repos Repositories
}

func New(repos Repositories) (factory *FactoryImp) {
	return &FactoryImp{repos}
}
