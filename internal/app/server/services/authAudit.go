package services

import (
	repositories "auth_audit/internal/app/repository/interfaces"
	"auth_audit/internal/app/repository/models"
)

const (
	authFailure = "auth_failure"
	authSuccess = "auth_success"
)

type LogAuthService struct {
	authAuditRepository repositories.AuthenticationAuditRepository
}

func NewLogAuth(authAuditRepository repositories.AuthenticationAuditRepository) *LogAuthService {
	return &LogAuthService{authAuditRepository: authAuditRepository}
}

func (s LogAuthService) LogAuthFailure(user *models.User) error {
	return s.createAuthEvent(authFailure, user)
}

func (s LogAuthService) LogAuthSuccess(user *models.User) error {
	return s.createAuthEvent(authSuccess, user)
}

func (s LogAuthService) createAuthEvent(actionType string, user *models.User) error {
	event := &models.AuthenticationAudit{
		UserID: user.ID,
		Event:  actionType,
	}

	if err := s.authAuditRepository.Save(event); err != nil {
		return err
	}

	return nil
}
