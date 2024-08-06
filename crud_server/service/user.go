package service

import (
	"crud-server/repository"
	"crud-server/types"
)

type User struct {
	userRepository *repository.UserRepository
}

func newUserService(userRpository *repository.UserRepository) *User {
	return &User{
		userRepository: userRpository,
	}
}

func (u *User) Create(newUser *types.User) error {
	return u.userRepository.Create(newUser)
}

func (u *User) Update(beofreUser *types.User, updatedUser *types.User) error {
	return u.userRepository.Update(beofreUser, updatedUser)
}

func (u *User) Delete(newUser *types.User) error {
	return u.userRepository.Delete(newUser)
}

func (u *User) Get() []*types.User {
	return u.userRepository.Get()
}
