package action

import (
	"github.com/fdsolutions/logan/metadata"
)

// Factory is compose of a set of methods to create actions
type Factory interface {
	MakeAction() LoganAction
}

// factoryImpl is a concrete factory
type factoryImpl struct{}

// New returns a action factory instance
func NewFactory() *factoryImpl {
	return new(factoryImpl)
}

// MakeActionFromMetadata creates a new action from a given metadata entry
func (f *factoryImpl) MakeActionFromMetadata(meta metadata.Entry) LoganAction {
	action := f.MakeAction()
	action.SetMetadata(meta)
	return action
}

// MakeAction creates a new action
func (f *factoryImpl) MakeAction() LoganAction {
	action := NewAction()
	return LoganAction(action)
}
