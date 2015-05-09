package agent

import (
	"github.com/fdsolutions/logan/action"
	"github.com/fdsolutions/logan/args"
	"github.com/fdsolutions/logan/metadata"
)

type API interface {
	PerformActionFromInput(input string) Status
	// RegisterRepo(r metadata.Repository) Status
	// LookupActionInRepos(goal string, repos []metadata.Repository) Status
	// Perform(action.LoganAction) Status
	// PrintOutput(output, printer io.Writer) Status
}

// Agent is a logan agent
type Agent struct {
	actionMaker action.Factory
	metaRepos   []metadata.Repository
	statusStack []Status
	parser      args.ParamParser
	API
}

// GetParser returns the agent's param parser
func (a *Agent) GetParser() args.ParamParser {
	return a.parser
}

// FromFactoryAndRepos returns a instnace of a agent with a given factory  and repositories.
func FromFactoryAndRepos(factory action.Factory, repos []metadata.Repository) (a *Agent) {
	a = &Agent{}
	a.actionMaker = factory
	a.metaRepos = repos
	return
}

// PerformActionFromInput processes user input and perform the action related ot the input.
// It follows a the template method pattern with a clear workflow
// You can change the behavior by overriding function from the agent API
func (a *Agent) PerformActionFromInput(input string) Status {
	s := a.parseUserInput(input)

	if s.GetCode() == StatusFail {
		return s
	}

	if _, ok := s.GetValue().(args.Arg); ok {
	}
	return NewStatus(StatusSuccess, nil)
}

func (a *Agent) parseUserInput(input string) Status {
	var s Status

	arg, err := args.ParseInputWithParser(input, a.GetParser())
	if err != nil {
		s = NewStatus(StatusFail, input)
		s.StackError(err)
		return s
	}

	s = NewStatus(StatusSuccess, arg)
	return s
}
