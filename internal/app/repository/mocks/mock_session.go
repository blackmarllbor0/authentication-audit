package mocks

import (
	"auth_audit/internal/app/repository/models"
	"github.com/stretchr/testify/mock"
)

type MockSessionRepository struct {
	mock.Mock
}

func (_m *MockSessionRepository) GetByToken(token string) (*models.Session, error) {
	args := _m.Called(token)
	return args.Get(0).(*models.Session), args.Error(1)
}

func (_m *MockSessionRepository) Create(session *models.Session) error {
	args := _m.Called(session)
	return args.Error(0)
}
