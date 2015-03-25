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

type CommandFactory struct{}

func NewFactory() (factory *CommandFactory) {
	return new(CommandFactory)
}

func (cf *CommandFactory) MakeCommandFromName(name string) Command {
	return NewCommandFromName(name)
}
