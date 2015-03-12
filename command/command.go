package command

import ()

type Expr string

type Arguments map[string]string


type Command interface {
	Metadata() Metadata
	GetArgs() Arguments
	Run()
}

type commandImp struct {
	meta Metadata
	args Arguments
	extras Arguments
}


func New(name string, args Arguments) (*CommandImp) {
	return &CommandImp{
		name
	}
}


