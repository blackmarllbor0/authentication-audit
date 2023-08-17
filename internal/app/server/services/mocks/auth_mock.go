package mocks

import (
	"auth_audit/internal/app/repository/models"
	"auth_audit/internal/app/server/DTO"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (_m *MockAuthService) ClearAuthAuditsByToken(token string) error {
	//TODO implement me
	panic("implement me")
}

func (_m *MockAuthService) GetAuthAuditByToken(token string) ([]DTO.AuthAuditDTO, error) {
	args := _m.Called(token)

	if DTOs := args.Get(0); DTOs != nil {
		return DTOs.([]DTO.AuthAuditDTO), nil
	}

	return nil, args.Error(1)
}

func (_m *MockAuthService) Register(dto DTO.RegisterUserDTO) (*models.Session, error) {
	args := _m.Called(dto)

	if s := args.Get(0); s != nil {
		return s.(*models.Session), nil
	}

	return nil, args.Error(1)
}

func (_m *MockAuthService) Login(dto DTO.LoginUserDTO) (*models.Session, error) {
	args := _m.Called(dto)

	if s := args.Get(0); s != nil {
		return s.(*models.Session), nil
	}

	return nil, args.Error(1)
}
