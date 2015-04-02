package agent

type Action struct {
}

// Agent perform actions related to user inputs
type Agent interface {
	Perform(Action)
}
