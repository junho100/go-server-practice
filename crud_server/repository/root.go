package repository

import (
	"sync"
)

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	User *UserRepository
}

func Newrepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			User: NewUserRepository(),
		}
	})

	return repositoryInstance
}
