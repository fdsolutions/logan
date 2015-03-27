package command

type Factory interface {
	MakeCommandFromMetatdata(meta Metadata) Command
}

type ConcreteFactory struct{}

func NewFactory() (factory *ConcreteFactory) {
	return new(ConcreteFactory)
}

// MakeCommandFromMetatdata creates a new command with the given metadata.
func (cf *ConcreteFactory) MakeCommandFromMetatdata(meta Metadata) (c Command) {
	c = NewCommand()
	c.SetMetadata(meta)
	return
}
