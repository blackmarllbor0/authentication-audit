package mocks

import (
	"auth_audit/internal/app/repository/models"
	"auth_audit/internal/app/server/DTO"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (_m *MockUserService) CreateUser(createUserDTO DTO.RegisterUserDTO) (*models.User, error) {
	args := _m.Called(createUserDTO)

	if user := args.Get(0); user != nil {
		return user.(*models.User), args.Error(1)
	}

	return nil, args.Error(1)
}

func (_m *MockUserService) GetUserByLogin(login string) (*models.User, error) {
	args := _m.Called(login)

	if user := args.Get(0); user != nil {
		return user.(*models.User), args.Error(1)
	}

	return nil, args.Error(1)
}
