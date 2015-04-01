package command

import (
	"github.com/fdsolutions/logan/errors"
	"github.com/fdsolutions/logan/metadata"
)

const (
	ErrInvalidCommand errors.ErrorCode = "Invalid command"
)

// Command is the base interface implemented by all commands
type Command interface {
	GetMetadata() metadata.Entry
	SetMetadata(meta metadata.Entry)
	Run() error
}

// ConcreteCommand is a concrete command
type Entry struct {
	meta metadata.Entry
}

func NewEntry() *Entry {
	return &Entry{}
}

// GetMetadata is the getter for meta attribute of the command
func (c *Entry) GetMetadata() metadata.Entry {
	return c.meta
}

// SetMetadata is the getter for meta attribute of the command
func (c *Entry) SetMetadata(meta metadata.Entry) {
	c.meta = meta
}

// Run executes the command
func (c *Entry) Run() error {
	return errors.New(ErrInvalidCommand)
}
