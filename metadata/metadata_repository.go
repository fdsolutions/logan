package metadata

import (
//"fmt"

//"github.com/fdsolutions/logan/log"
)

// Repository is a set of methods that a metadata repository must implement
// TODO: Add 'error' has a second value returned by the interface functions
type Repository interface {
	FindAll() []Entry
	FindByGoal(string) (Entry, bool)
	FindByContext(string) []Entry
	SetStore(Store)
}

// repositoryImpl is a repository used to find command metadata
// It has its own store.
type repositoryImpl struct {
	store Store
}

// NewFromStore returns a new repository and set its store at the same time.
func NewRepositoryFromStore(s Store) Repository {
	repo := NewRepository()
	repo.SetStore(s)
	return repo
}

// NewRepository returns a repository instance
func NewRepository() Repository {
	return Repository(&repositoryImpl{})
}

// SetStore set a the repository store
func (r *repositoryImpl) SetStore(s Store) {
	r.store = s
}

// GetStore returns the metadata store of the repository
func (r *repositoryImpl) getStore() Store {
	return r.store
}

// FindAll retrieves all metadata entries from the repository
func (r *repositoryImpl) FindAll() (entries []Entry) {
	entries, _ = r.getStore().QueryAll()
	return
}

// FindByGoal get the first entry of the given goal even if the store
// retrieve more than one entries for that goal.
// TODO: Find a way to Handle conflicts when multiple entries are retrived
func (r *repositoryImpl) FindByGoal(g string) (entry Entry, found bool) {
	entries := r.getStore().Query(PredicateForGoal(g))
	if len(entries) < 1 {
		found = false
		return
	}
	return entries[0], true
}

// FindByContext returns all metadata entries related to a given context
func (r *repositoryImpl) FindByContext(ctx string) (entries []Entry) {
	entries = r.getStore().Query(PredicateForContext(ctx))
	return
}
