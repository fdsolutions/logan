package command

// Repository is the interface that every repository
// must implement
type Repository interface {
	FindCommandByName(string) (Command, error)
}

// Repositories is a collection of repository
type Repositories []Repository
