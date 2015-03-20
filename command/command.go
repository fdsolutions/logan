package command

import (
	"github.com/fdsolutions/logan/command/metadata"
)

const (
	commandNamePartSeparator string = ":"
	cumberOfCommandNameParts        = "3"
)

// Command is the base interface implemented by all commands
type Command interface {
	Metadata() Metadata
	Run() error
}

// Imp is a concrete command object
type Imp struct {
	meta Metadata
}

// New is the command construtor
func New(name string) *Imp {
	meta := metadata.New()
	return &Imp{}
}
