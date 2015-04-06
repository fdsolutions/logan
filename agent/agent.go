package agent

import (
	"io"

	"github.com/fdsolutions/logan/action"
	"github.com/fdsolutions/logan/errors"
	"github.com/fdsolutions/logan/metadata"
)

const (
	ErrFailToPerfomAction errors.ErrorCode = "Fail to perfom action"
)

type Api interface {
	ParseUserInput(input string) Status
	RegisterRepo(r metadata.Repository) Status
	LookupActionInRepos(goal string, repos []metadata.Repository) Status
	Perform(action.LoganAction) Status
	PrintOutput(output, printer io.Writer) Status
}

// Agent is a logan agent
type Agent struct {
	actionMaker action.Factory
	metaRepos   []metadata.Repository
	statusStack []Status
	Api
}

func FromFactoryAndRepos(f action.Factory, repos []metadata.Repository) (a *Agent) {
	a = &Agent{}
	a.actionMaker = f
	a.metaRepos = repos
	return
}

// PerformActionFromInput processes user input and perform the action related ot the input.
// It follows a the template method pattern with a clear workflow
// You can change the behavior by overriding function from the agent API
func (a *Agent) PerformActionFromInput(input string) Status {
	// var (
	// 	status  Status
	// 	action  action.LoganAction
	// 	outFile io.Writer
	// )
	return NewStatus(StatusSuccess, nil)
}
