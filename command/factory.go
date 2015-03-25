package command

import (
	errors "github.com/fdsolutions/logan/errors"
)

const (
	InvalidCommandName errors.ErrorCode = "Invalid command name"
)

type Factory interface {
	MakeCommandFromName(name string) Command
	MakeCommandFromMetadata(meta Metadata) Command
}

type ConcreteFactory struct{}

func NewFactory() (factory *ConcreteFactory) {
	return new(ConcreteFactory)
}

func (cf *ConcreteFactory) MakeCommandFromName(name string) Command {
	return NewCommandFromName(name)
}
