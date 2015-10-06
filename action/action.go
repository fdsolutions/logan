package action

import (
	"os/exec"

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
	BuildCommand() *exec.Cmd
}

// actionImpl is a concrete action
type actionImpl struct {
	meta *metadata.Entry
}

func NewAction() *actionImpl {
	return &actionImpl{}
}

// GetMetadata is the getter for meta attribute of the action
func (a *actionImpl) GetMetadata() metadata.Entry {
	return *a.meta
}

// SetMetadata is the getter for meta attribute of the action
func (a *actionImpl) SetMetadata(meta metadata.Entry) {
	a.meta = meta.Clone()
}

func (a *actionImpl) BuildCommand() *exec.Cmd {
	return new(exec.Cmd)
}
