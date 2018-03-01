package mock

import (
	"home-collect/domain"

	"github.com/stretchr/testify/mock"
)

type IUserRepository struct {
	mock.Mock
}

func (mu *IUserRepository) InsertUser(user domain.User) error {
	args := mu.Called(user)
	return args.Error(0)
}
func (mu *IUserRepository) FindAllUser() ([]domain.User, error) {
	args := mu.Called()
	var users []domain.User
	return users, args.Error(1)
}
func (mu *IUserRepository) FindByIdUser(id string) (domain.User, error) {
	args := mu.Called(id)
	var user domain.User
	return user, args.Error(1)
}
func (mu *IUserRepository) DeleteUser(user domain.User) error {
	args := mu.Called(user)
	return args.Error(0)
}
func (mu *IUserRepository) UpdateUser(user domain.User) error {
	args := mu.Called(user)
	return args.Error(0)
}
