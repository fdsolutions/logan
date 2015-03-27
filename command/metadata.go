package command

import (
	"strings"

	errors "github.com/fdsolutions/logan/errors"
)

const (
	ErrConflictingCommandFound errors.ErrorCode = `Command conflict : 
at least two command found for the same name`
	ErrInvalidFilePath errors.ErrorCode = "Invalid store source file path"
)

// Metadata holds command specification
type Metadata struct {
	Name    string
	Intent  string
	Target  string
	Context string
}

// NewMetadataFromName instanciate a new metadata from command name
func NewMetadataFromName(name string) Metadata {
	intent, target, ctx := GetCommandNameParts(name)
	return NewMetadata(intent, target, ctx)
}

// New create a metadata object from its properties (intent, target, context).
// Make sure parameters are in the right order.
func NewMetadata(intent string, target string, context string) Metadata {
	name := strings.Join([]string{intent, target, context}, commandNamePartSeparator)
	return Metadata{name, intent, target, context}
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
	FindByIntent(intent string) []Metadata
	FindByTarget(target string) []Metadata
	FindByContext(ctx string) []Metadata
}

// FileMetadataStore is a metadata store using file as a data source
type FileMetadataStore struct {
	FilePath string
}

func NewFileMetadataStore(path string) (fs *FileMetadataStore, err error) {
	if path == "" {
		err = errors.New(ErrInvalidFilePath)
		return
	}
	fs = &FileMetadataStore{path}
	return
}

func (fs *FileMetadataStore) FindByName(name string) (meta Metadata, err error) {
	return
}
