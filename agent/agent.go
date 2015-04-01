package agent

type Action struct {
}

// Agent is the one which execute command
type Agent interface {
	Perform(Action)
}
