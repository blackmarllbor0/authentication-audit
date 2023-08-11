package mocks

import "github.com/stretchr/testify/mock"

type MockAuthAudit struct {
	mock.Mock
}

func (_m *MockAuthAudit) Create(event string, userID uint) error {
	args := _m.Called(event, userID)
	return args.Error(0)
}
