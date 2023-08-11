package mocks

import (
	"auth_audit/internal/app/repository/models"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (_m *MockUserRepository) Create(user *models.User) error {
	args := _m.Called(user)
	return args.Error(0)
}

func (_m *MockUserRepository) GetByLogin(login string) (*models.User, error) {
	args := _m.Called(login)

	if user := args.Get(0); user != nil {
		return user.(*models.User), args.Error(1)
	}

	return nil, args.Error(1)
}

func (_m *MockUserRepository) GetById(ID uint) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (_m *MockUserRepository) Block(userID uint) error {
	panic("")
}

func (_m *MockUserRepository) IncrementFailedLoginAttempts(userID uint) (attempt int, err error) {
	panic("")
}
