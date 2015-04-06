package action

import (
	"github.com/fdsolutions/logan/errors"
	"github.com/fdsolutions/logan/metadata"
)

const (
	ErrInvalidAction errors.ErrorCode = "Invalid action"
)

// LoganAction is the base interface implemented by all actions
type LoganAction interface {
	GetMetadata() metadata.Entry
	SetMetadata(meta metadata.Entry)
	Run() error
}

// actionImpl is a concrete action
type actionImpl struct {
	meta metadata.Entry
}

func NewAction() *actionImpl {
	return &actionImpl{}
}

// GetMetadata is the getter for meta attribute of the action
func (c *actionImpl) GetMetadata() metadata.Entry {
	return c.meta
}

// SetMetadata is the getter for meta attribute of the action
func (c *actionImpl) SetMetadata(meta metadata.Entry) {
	c.meta = meta
}

// Run executes the action
func (c *actionImpl) Run() error {
	return errors.New(ErrInvalidAction)
}
