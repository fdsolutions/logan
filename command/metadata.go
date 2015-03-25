package command

// Metadata holds command specification
type Metadata struct {
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
	return Metadata{intent, target, context}
}

// MetadataStore defines abstract operations to implement for a metadata store
type MetadataStore interface {
	GetByName(name string) *Metadata
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
func (r *MetadataRepository) FindCommandByName(name string) (Command, error) {
	var s MetadataStore = r.GetMetadataStore()
	var f Factory = r.GetFactory()
	// Get metadata related to the command name
	meta := s.GetByName(name)
	return f.MakeCommandFromMetatdata(meta), nil
}

// GetFactory is used as a getter for factory property of repository
func (r *MetadataRepository) GetFactory() Factory {
	return r.factory
}

// GetStore returns the metadata store of the repository
func (r *MetadataRepository) GetMetadataStore() MetadataStore {
	return r.store
}
