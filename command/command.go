package command

import (
	"github.com/fdsolutions/logan/errors"
)

const (
	commandNamePartSeparator string = ":"
	numberOfCommandNameParts int    = 3
	noPartValue              string = ""

	ErrInvalidCommand errors.ErrorCode = "Invalid command"
)

// Command is the base interface implemented by all commands
type Command interface {
	GetMetadata() Metadata
	SetMetadata(meta Metadata)
	Run() error
}

// ConcreteCommand is a concrete command
type ConcreteCommand struct {
	meta Metadata
}

func NewCommand() *ConcreteCommand {
	return &ConcreteCommand{}
}

// Metadata is the getter for meta attribute of the command
func (c *ConcreteCommand) GetMetadata() Metadata {
	return c.meta
}

// Metadata is the getter for meta attribute of the command
func (c *ConcreteCommand) SetMetadata(meta Metadata) {
	c.meta = meta
}

func (c *ConcreteCommand) Run() error {
	return errors.New(ErrInvalidCommand)
}
