package agent

import (
	"github.com/fdsolutions/logan/action"
	"github.com/fdsolutions/logan/args"
	"github.com/fdsolutions/logan/errors"
	"github.com/fdsolutions/logan/metadata"
)

type API interface {
	ParseUserInput(input string) Status
	LookupActionInRepos(goal string, repos []metadata.Repository) Status
	//PerformActionFromInput(input string) Status
	// RegisterRepo(r metadata.Repository) Status
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
func (ag *Agent) GetParser() args.ParamParser {
	return ag.parser
}

// FromFactoryAndRepos returns a instnace of a agent with a given factory  and repositories.
func FromFactoryAndRepos(factory action.Factory, repos []metadata.Repository) *Agent {
	ag := &Agent{}
	ag.actionMaker = factory
	ag.metaRepos = repos
	return ag
}

// PerformActionFromInput processes user input and perform the action related ot the input.
// It follows a the template method pattern with a clear workflow
// You can change the behavior by overriding function from the agent API
// func (a *Agent) PerformActionFromInput(input string) Status {
// 	s := a.ParseUserInput(input)

// 	if s.GetCode() == StatusFail {
// 		return s
// 	}

// 	if _, ok := s.GetValue().(args.Arg); ok {
// 	}
// 	return NewStatus(StatusSuccess, nil)
// }

func (ag *Agent) ParseUserInput(input string) (s Status) {
	arg, err := args.ParseInputWithParser(input, ag.GetParser())
	if err != nil {
		s = NewStatus(StatusFail, input)
		s.StackError(errors.New(errors.ErrInvalidUserInput))
		return s
	}
	return NewStatus(StatusSuccess, arg)
}

func (ag *Agent) LookupActionInRepos(goal string, repos []metadata.Repository) Status {
	if goal == "" || repos == nil {
		s := NewStatus(StatusFail, []interface{}{goal, repos})
		s.StackError(errors.New(errors.ErrInvalidGoal))
		return s
	}

	a := ag.pickFirstActionInReposMatchingGoal(goal, repos)
	return NewStatus(StatusSuccess, a)
}

func (ag *Agent) pickFirstActionInReposMatchingGoal(g string, repos []metadata.Repository) (a action.LoganAction) {
	var meta metadata.Entry
	for _, r := range repos {
		m, found := r.FindByGoal(g)
		if found {
			meta = m
			break
		}
	}
	return ag.actionMaker.MakeActionFromMetadata(meta)
}
