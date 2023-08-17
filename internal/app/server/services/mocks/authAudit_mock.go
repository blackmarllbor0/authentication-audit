package mocks

import (
	"auth_audit/internal/app/repository/models"
	"github.com/stretchr/testify/mock"
)

type MockAuthAudit struct {
	mock.Mock
}

func (_m *MockAuthAudit) ClearAuthAuditsByToken(userID uint) error {
	args := _m.Called(userID)
	return args.Error(0)
}

func (_m *MockAuthAudit) GetAllAuditsByUserID(userID uint) ([]models.AuthenticationAudit, error) {
	args := _m.Called(userID)
	return args.Get(0).([]models.AuthenticationAudit), args.Error(1)
}

func (_m *MockAuthAudit) Create(event string, userID uint) error {
	args := _m.Called(event, userID)
	return args.Error(0)
}
