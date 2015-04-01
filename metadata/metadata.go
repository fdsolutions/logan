package metadata

import (
	"strings"

	"github.com/fdsolutions/logan/errors"
)

const (
	goalPartsSeparator string = ":"
	numberOfGoalParts  int    = 3
	noPartValue        string = ""

	ErrConflictingCommandFound errors.ErrorCode = `Command conflict : 
at least two command found for the same goal`
	ErrInvalidFilePath       errors.ErrorCode = "Invalid store source file path"
	ErrUnsupportedFileFormat errors.ErrorCode = "Unsupported file format: can't load data from store file"
)

// Entry holds command metadata
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

// NewFromGoal instanciate a new metadata from command goal
func NewFromGoal(goal string) *Entry {
	intent, target, ctx := SplitInGoalParts(goal)
	return NewFromGoalParts(intent, target, ctx)
}

// NewFromGoalParts creates a metadata object from its goal.
// A goal is compose of (intent, target, context).
// Make sure parameters are in the right order.
func NewFromGoalParts(intent string, target string, context string) *Entry {
	var entry *Entry
	goal := strings.Join([]string{intent, target, context}, goalPartsSeparator)
	entry = NewEntry()
	entry.Goal = goal
	entry.Intent = intent
	entry.Target = target
	entry.Context = context
	return entry
}

// Repository is a set of methods that a metadata repository must implement
type Repository interface {
	GetStore() Store
}

// repositoryImpl is a repository used to find command metadata
// It has its own store.
type repositoryImpl struct {
	store Store
}

// NewRepository returns a new repository
func NewRepository() *repositoryImpl {
	return &repositoryImpl{}
}

// NewFromStore returns a new repository and set its store at the same time.
func NewRepositoryFromStore(s Store) *repositoryImpl {
	repo := NewRepository()
	repo.store = s
	return repo
}

// GetStore returns the metadata store of the repository
func (r *repositoryImpl) GetStore() Store {
	return r.store
}
