package command

import (
	"github.com/fdsolutions/logan/errors"
)

const (
	commandNamePartSeparator string = ":"
	numberOfCommandNameParts        = "3"
	noPartValue                     = ""

	InvalidCommand errors.ErrorCode = "Invalid command"
)

// Command is the base interface implemented by all commands
type Command interface {
	Metadata() Metadata
	Run() error
}

// Imp is a concrete command object
type ConcreteCommand struct {
	meta Metadata
}

// New is the command construtor
func NewCommandFromName(name string) *ConcreteCommand {
	meta := NewMetadataFromName(name)
	return &ConcreteCommand{meta}
}

// Metadata is the getter for meta attribute of the command
func (c *ConcreteCommand) Metadata() Metadata {
	return c.meta
}

func (c *ConcreteCommand) Run() error {
	return errors.New(InvalidCommand)
}
