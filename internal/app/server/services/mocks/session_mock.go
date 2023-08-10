package mocks

import (
	"auth_audit/internal/app/repository/models"
	"github.com/stretchr/testify/mock"
)

type MockSessionService struct {
	mock.Mock
}

func (_m *MockSessionService) Create(userID uint) (*models.Session, error) {
	args := _m.Called(userID)

	if s := args.Get(0); s != nil {
		return s.(*models.Session), args.Error(1)
	}

	return nil, args.Error(1)
}
