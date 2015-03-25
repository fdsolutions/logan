package command

// Repository is a common interface for all repos
type Repository interface {
	FindCommandByName(string) (Command, error)
}

// Repositories is a collection of repository
type Repositories []Repository
