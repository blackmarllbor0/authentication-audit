package mocks

import (
	"auth_audit/internal/app/repository/models"
	"github.com/stretchr/testify/mock"
)

type MockSessionRepository struct {
	mock.Mock
}

func (_m *MockSessionRepository) Create(session *models.Session) error {
	args := _m.Called(session)
	return args.Error(0)
}
