package metadata

import (
	"encoding/json"
	"io/ioutil"

	"github.com/fdsolutions/logan/errors"
	"github.com/fdsolutions/logan/helper"
)

// Store defines abstract operations to implement for a metadata store
type Store interface {
	Filepath() string
	QueryAll() []Entry
	Query(cond func(Entry) bool) []Entry
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

// Query tries to find metadata entries that match the given predicace
func (fs *FileStore) Query(cond func(Entry) bool) []Entry {
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
	err = fs.load()
	entries = fs.data
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

func (fs *FileStore) filterThrough(cond func(Entry) bool) []Entry {
	selectedEntries := []Entry{}
	for _, entry := range fs.data {
		if ok := cond(entry); ok {
			selectedEntries = append(selectedEntries, entry)
		}
	}
	return selectedEntries
}
