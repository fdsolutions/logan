package command

// RepositoryReader defines reading operations.
type RepositoryReader interface {
	Query(cond map[string]string) []Metadata
	GetByName(name string) Metadata
}

// RepositoryReadWriter defines writing operations.
type RepositoryReadWriter interface {
	RepositoryReader
	Put(Metadata) error
	Del() (Metadata, error)
}

// ReadOnlyMetadataStore allows read operations from the repository.
type ReadOnlyMetadataStore interface {
	RepositoryReader
}

// MeatadataStore allows all operations from repository.
type MeatadataStore interface {
	RepositoryReadWriter
}

// Repositories rpresent collection of repository
type Repositories []Repository

// Repository is a common interface for all repos
type Repository interface {
	FindByName(string) (Metadata, error)
}

// RepositoryImp is a concrete repository.
// It has its own store.
type RepositoryImp struct {
	store MeatadataStore
}

// NewRepositoryWithStore returns a new repository and set its store.
func NewRepositoryWithStore(store MeatadataStore) *RepositoryImp {
	repo := NewRepository()
	repo.store = store
	return repo
}

// NewRepository returns a new repository
func NewRepository() *RepositoryImp {
	return &RepositoryImp{}
}

// FindCmdByName gets command metadata by command name.
func (r *RepositoryImp) FindCmdByName(name string) Command {

}
