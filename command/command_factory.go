package command

// Factory is compose of a set of methods to create commands
type Factory interface {
	MakeCommand() Command
}

// factoryImpl is a concrete factory
type factoryImpl struct{}

// New returns a command factory instance
func NewFactory() *factoryImpl {
	return new(factoryImpl)
}

// MakeCommand creates a new command
func (f *factoryImpl) MakeCommand() Command {
	return Command(NewEntry())
}
