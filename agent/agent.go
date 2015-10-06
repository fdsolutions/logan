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
	PerformAction(action.LoganAction) Status
	// RegisterRepo(r metadata.Repository) Status
	//PerformActionFromInput(input string) Status
	// PrintOutput(output, printer io.Writer) Status
}

// Agent is a logan agent
type Agent struct {
	actionMaker   action.Factory
	statusStack   []Status
	parser        args.ParamParser
	metadataRepos []metadata.Repository
	Output        string

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

// PerformAction execute the command related to the given action
func (a *Agent) PerformAction(act action.LoganAction) Status {

	return NewStatus(StatusSuccess, nil)
}

// ParseUserInput parse the user input and return the related argument's values
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
