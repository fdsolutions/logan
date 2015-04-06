package metadata

import (
	"encoding/json"
	"io/ioutil"

	"github.com/fdsolutions/logan/errors"
	"github.com/fdsolutions/logan/helper"
)

// Predicate is a closure that identicate whether or not an entry
// be added to a query result
type Predicate func(ent Entry) bool

// Store defines abstract operations to implement for a metadata store
type Store interface {
	Filepath() string
	QueryAll() ([]Entry, error)
	Query(cond Predicate) []Entry
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

// PredicateForGoal create a predicate to retrieve metadata entry
// of the given goal.
func PredicateForGoal(goal string) Predicate {
	var predicate Predicate = func(entry Entry) bool {
		return (entry.Goal == goal)
	}
	return predicate
}

// PredicateForContext create a predicate to retrieve metadata entry
// of the given context.
func PredicateForContext(ctx string) Predicate {
	var predicate Predicate = func(entry Entry) bool {
		_, _, context := SplitInGoalParts(entry.Goal)
		return (context == ctx)
	}
	return predicate
}

// Query tries to find metadata entries that match the given predicace
func (fs *FileStore) Query(cond Predicate) []Entry {
	entries, _ := fs.QueryAll()
	if cond == nil {
		return entries
	}
	return fs.filterThrough(cond)
}

// QueryAll returns all metadata contains in the store
// It returns an error as a second return value if something bad happens,
// this to enable error handling
func (fs *FileStore) QueryAll() (entries []Entry, err error) {
	if fs.HasDataAlreadyLoaded() {
		return fs.data, nil
	}
	err = fs.load()
	return fs.data, err
}

// HasDataAlreadyLoaded checks whether or not data are loaded
func (fs *FileStore) HasDataAlreadyLoaded() bool {
	return (fs.data != nil && len(fs.data) > 1)
}

func (fs *FileStore) load() (err error) {
	jsonContentAsBytes, _ := ioutil.ReadFile(fs.Filepath())
	err = fs.loadFromJSON(jsonContentAsBytes)
	return
}

func (fs *FileStore) loadFromJSON(content []byte) (err error) {
	var entries []*Entry
	err = json.Unmarshal(content, &entries)
	if err != nil {
		return errors.New(ErrUnsupportedFileFormat)
	}
	fs.setDataAndAutoFillEachEntry(entries)
	return
}

func (fs *FileStore) setDataAndAutoFillEachEntry(entries []*Entry) {
	fs.data = []Entry{}
	for _, entry := range entries {
		entry.AutoFill()
		fs.data = append(fs.data, *entry)
	}
}

func (fs *FileStore) filterThrough(cond Predicate) []Entry {
	filteredEntries := []Entry{}
	for _, entry := range fs.data {
		if ok := cond(entry); ok {
			filteredEntries = append(filteredEntries, entry)
		}
	}
	return filteredEntries
}
