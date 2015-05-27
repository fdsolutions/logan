package agent

import (
	"github.com/fdsolutions/logan/action"
	"github.com/fdsolutions/logan/args"
	"github.com/fdsolutions/logan/errors"
	//"github.com/fdsolutions/logan/helper"
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
	actionMaker   action.Factory
	statusStack   []Status
	parser        args.ParamParser
	metadataRepos []metadata.Repository

	API
}

// FromFactoryAndRepos returns a instnace of a agent with a given factory  and repositories.
func FromFactoryAndRepos(factory action.Factory, repos []metadata.Repository) (ag *Agent, err errors.LoganError) {
	if factory == nil {
		err = errors.New(errors.ErrMissingActionFactory)
		return
	}
	ag = &Agent{}
	ag.actionMaker = factory
	ag.metadataRepos = repos
	return
}

// GetParser returns the agent's param parser
func (ag *Agent) GetParser() args.ParamParser {
	return ag.parser
}

// GetMetadataRepos returns the list of repos register on the agent
func (ag *Agent) GetMetadataRepos() []metadata.Repository {
	return ag.metadataRepos
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
	var (
		s Status
		a action.LoganAction
	)

	if goal == "" || repos == nil {
		s = NewStatus(StatusFail, []interface{}{goal, repos})
		s.StackError(errors.New(errors.ErrInvalidGoal))
		return s
	}

	a, err := ag.pickFirstActionInReposMatchingGoal(goal, repos)

	if err != nil {

		//helper.W("maker", a)
		s = NewStatus(StatusFail, []interface{}{goal, repos})
		s.StackError(errors.New(errors.ErrActionNotFound))
		return s
	}

	return NewStatus(StatusSuccess, a)
}

// TODO : Change the way I look for the metadata
func (ag *Agent) pickFirstActionInReposMatchingGoal(g string, repos []metadata.Repository) (a action.LoganAction, err errors.LoganError) {
	var (
		meta  metadata.Entry
		found bool
	)

	// If there a action factory present on the agent
	if ag.actionMaker == nil {
		err = errors.New(errors.ErrMissingActionFactory)
		return
	}

	for _, r := range repos {
		if meta, found = r.FindByGoal(g); found {
			break
		}
	}

	if !found {
		err = errors.New(errors.ErrNoMetadataFound)
		return
	}

	a = ag.actionMaker.MakeActionFromMetadata(meta)
	return
}
