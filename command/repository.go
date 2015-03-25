package command

// Store defines writing operations.
// - Put(Metadata) error
// - Query(conditions map[string]string) []Metadata
type Store interface {
	Get(key string) Metadata
}

// MetadataStore allows all operations from repository.
type MetadataStore interface {
	Store
}

// Repository is a common interface for all repos
type Repository interface {
	FindCommandByName(string) (Command, error)
}

// Repositories is a collection of repository
type Repositories []Repository

// ConcreteRepository is a concrete repository.
// It has its own store.
type ConcreteRepository struct {
	store MetadataStore
}

// NewRepositoryWithStore returns a new repository and set its store.
func NewRepositoryWithStore(store MetadataStore) Repository {
	repo := NewRepository()
	repo.store = store
	return repo
}

// NewRepository returns a new repository
func NewRepository() *ConcreteRepository {
	return &ConcreteRepository{}
}

// FindCmdByName gets command metadata by command name.
func (r *ConcreteRepository) FindCommandByName(name string) (Command, error) {
	return NewCommandFromName(name), nil
}
