package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	errors "github.com/fdsolutions/logan/errors"
	helpers "github.com/fdsolutions/logan/helpers"
)

const (
	ErrConflictingCommandFound errors.ErrorCode = `Command conflict : 
at least two command found for the same name`
	ErrInvalidFilePath       errors.ErrorCode = "Invalid store source file path"
	ErrUnsupportedFileFormat errors.ErrorCode = "Unsupported file format: can't load data from store file"
)

// Metadata holds command specification
type Metadata struct {
	Name           string   `json:"name"`
	Intent         string   `json:"intent"`
	Target         string   `json:"target"`
	Context        string   `json:"context"`
	Path           string   `json:"path"`
	RequiredParams []string `json:"required_params"`
}

// NewMetadataFromName instanciate a new metadata from command name
func NewMetadataFromName(name string) *Metadata {
	intent, target, ctx := GetCommandNameParts(name)
	return NewMetadataFromGoal(intent, target, ctx)
}

// New create a metadata object from its goal.
// A goal is compose of (intent, target, context).
// Make sure parameters are in the right order.
func NewMetadataFromGoal(intent string, target string, context string) *Metadata {
	var m *Metadata
	name := strings.Join([]string{intent, target, context}, commandNamePartSeparator)
	m = NewMetadata()
	m.Name = name
	m.Intent = intent
	m.Target = target
	m.Context = context
	return m
}

func NewMetadata() *Metadata {
	return &Metadata{}
}

// MetadataRepository is a repository used to find command metadata
// It has its own store.
type MetadataRepository struct {
	store   MetadataStore
	factory Factory
}

// NewRepository returns a new repository
func NewMetadataRepository() *MetadataRepository {
	return &MetadataRepository{}
}

// NewRepositoryWithStore returns a new repository and set its store.
func NewFromStoreAndFactory(s MetadataStore, f Factory) Repository {
	repo := NewMetadataRepository()
	repo.store = s
	repo.factory = f
	return repo
}

// FindCommandByName looks for command metadata in the repository by the given name and
// create a new command instance from the metadata found.
func (r *MetadataRepository) FindCommandByName(name string) (c Command, err error) {
	s := r.GetMetadataStore()
	f := r.GetFactory()

	// Get metadata related to the command name
	meta, err := s.FindByName(name)
	if err != nil {
		return
	}
	return f.MakeCommandFromMetatdata(meta), nil
}

// GetStore returns the metadata store of the repository
func (r *MetadataRepository) GetMetadataStore() MetadataStore {
	return r.store
}

// GetFactory is used as a getter for factory property of repository
func (r *MetadataRepository) GetFactory() Factory {
	return r.factory
}

// MetadataStore defines abstract operations to implement for a metadata store
type MetadataStore interface {
	FindByName(name string) (Metadata, error)
	FindAll() []Metadata
	// FindByIntent(intent string) []Metadata
	// FindByTarget(target string) []Metadata
	// FindByContext(ctx string) []Metadata
	Filepath() string
}

// FileMetadataStore is a metadata store using file as a data source
type FileMetadataStore struct {
	filePath string
	data     []Metadata
}

// NewFileMetadataStore returns an instance of a FileMetadataStore
func NewFileMetadataStore(path string) (fs *FileMetadataStore, err error) {
	if path == "" || helpers.FileDoesntExist(path) {
		err = errors.New(ErrInvalidFilePath)
		return
	}
	fs = &FileMetadataStore{path, nil}
	return
}

// Filepath returns a source file path of the metadata store
func (fs *FileMetadataStore) Filepath() string {
	return fs.filePath
}

// FindAll returns all metadata contains in the store
// It returns an error as a second return value if something bad happens,
// this to enable error handling
func (fs *FileMetadataStore) FindAll() (metas []Metadata, err error) {
	err = fs.load()
	metas = fs.data
	return
}

func (fs *FileMetadataStore) load() (err error) {
	jsonContentAsBytes, _ := ioutil.ReadFile(fs.filePath)
	err = fs.loadFromJSON(jsonContentAsBytes)
	return
}

func (fs *FileMetadataStore) loadFromJSON(content []byte) (err error) {
	var metas []Metadata
	err = json.Unmarshal(content, &metas)
	helpers.W("loadFromJSON: ", fmt.Sprintf("%v, err: %v", metas, err))
	if err != nil {
		return errors.New(ErrUnsupportedFileFormat)
	}
	fs.data = metas
	return
}
