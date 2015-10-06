package metadata

import (
	"strings"

	"github.com/fdsolutions/logan/errors"
)

const (
	goalPartsSeparator string = ":"
	numberOfGoalParts  int    = 3
	noPartValue        string = ""

	ErrConflictingActionFound errors.ErrorCode = `Action conflict : 
at least two actions found for the same goal`
	ErrInvalidFilePath       errors.ErrorCode = "Invalid store source file path"
	ErrUnsupportedFileFormat errors.ErrorCode = "Unsupported file format: can't load data from store file"
)

// Entry holds action's metadata
type Entry struct {
	Goal           string   `json:"goal"`
	Intent         string   `json:"intent"`
	Target         string   `json:"target"`
	Context        string   `json:"context"`
	Path           string   `json:"path"`
	RequiredParams []string `json:"required_params"`
}

// New is the matadata entry constructor
func NewEntry() *Entry {
	return &Entry{}
}

// FromGoal instanciate a new metadata entry from an action goal
func FromGoal(goal string) *Entry {
	intent, target, ctx := SplitInGoalParts(goal)
	return FromGoalParts(intent, target, ctx)
}

// FromGoalParts creates a metadata object from its goal.
// A goal is compose of (intent, target, context).
// Make sure parameters are in the right order.
func FromGoalParts(intent string, target string, context string) *Entry {
	var entry *Entry
	goal := strings.Join([]string{intent, target, context}, goalPartsSeparator)
	entry = NewEntry()
	// Removes possible ':' around the goal name.
	entry.Goal = strings.Trim(goal, goalPartsSeparator)
	entry.Intent = intent
	entry.Target = target
	entry.Context = context
	return entry
}

// AutoFill make sure that fiels Intent, Target, Context
// are filled with values got from the Goal field
func (entry *Entry) AutoFill() {
	intent, target, context := SplitInGoalParts(entry.Goal)
	entry.Intent = intent
	entry.Target = target
	entry.Context = context
}

// SetPath set the path of the metadata entry
func (entry *Entry) SetPath(path string) {
	entry.Path = path
}

// SetRequiredParams set the required params of the metadata entry
func (entry *Entry) SetRequiredParams(params []string) {
	entry.RequiredParams = params
}

// Clone returns a copy of the given metadata
func (entry *Entry) Clone() *Entry {
	clone := FromGoal(entry.Goal)
	clone.SetPath(entry.Path)
	clone.SetRequiredParams(entry.RequiredParams)
	return clone
}
