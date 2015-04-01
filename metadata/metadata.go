package metadata

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/fdsolutions/logan/errors"
	"github.com/fdsolutions/logan/helper"
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

// Store defines abstract operations to implement for a metadata store
type Store interface {
	Filepath() string
	QueryAll() []Entry
}

// FileStore is a metadata store using file as a data source
type FileStore struct {
	filePath string
	data     []Entry
}

// NewFileStore returns an instance of a FileStore
func NewFileStore(path string) (fs *FileStore, err error) {
	if path == "" || helper.FileDoesntExist(path) {
		err = errors.New(ErrInvalidFilePath)
		return
	}
	fs = &FileStore{path, nil}
	return
}

// Filepath returns a source file path of the metadata store
func (fs *FileStore) Filepath() string {
	return fs.filePath
}

// QueryAll returns all metadata contains in the store
// It returns an error as a second return value if something bad happens,
// this to enable error handling
func (fs *FileStore) QueryAll() (metas []Entry, err error) {
	err = fs.load()
	metas = fs.data
	return
}

func (fs *FileStore) load() (err error) {
	jsonContentAsBytes, _ := ioutil.ReadFile(fs.filePath)
	err = fs.loadFromJSON(jsonContentAsBytes)
	return
}

func (fs *FileStore) loadFromJSON(content []byte) (err error) {
	var metas []Entry
	err = json.Unmarshal(content, &metas)
	if err != nil {
		return errors.New(ErrUnsupportedFileFormat)
	}
	fs.data = metas
	return
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
