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

	if s := args.Get(0); s != nil {
		return s.(*models.Session), nil
	}

	return nil, args.Error(1)
}

func (_m *MockSessionRepository) Create(session *models.Session) error {
	args := _m.Called(session)
	return args.Error(0)
}
