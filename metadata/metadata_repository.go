package metadata

// Repository is a set of methods that a metadata repository must implement
type Repository interface {
	FindAll() []Entry
	FindByGoal(string) Entry
	FindByContext(string) []Entry
}

// repositoryImpl is a repository used to find command metadata
// It has its own store.
type repositoryImpl struct {
	store Store
}

// NewFromStore returns a new repository and set its store at the same time.
func NewRepositoryFromStore(s Store) *repositoryImpl {
	repo := NewRepository()
	repo.SetStore(s)
	return repo
}

// NewRepository returns a repository instance
func NewRepository() *repositoryImpl {
	return &repositoryImpl{}
}

// SetStore set a the repository store
func (r *repositoryImpl) SetStore(s Store) {
	r.store = s
}

// GetStore returns the metadata store of the repository
func (r *repositoryImpl) getStore() Store {
	return r.store
}

func (r *repositoryImpl) FindAll() (entries []Entry) {
	return
}

func (r *repositoryImpl) FindByGoal(g string) (entry Entry) {
	return
}

func (r *repositoryImpl) FindByContext(ctx string) (entries []Entry) {
	return
}
