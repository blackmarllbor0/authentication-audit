package services

import (
	"auth_audit/internal/app/repository/interfaces"
	"auth_audit/internal/app/repository/models"
	"time"
)

const (
	SuccessfulLogin   = "successful login"
	IncorrectPassword = "incorrect password"
	Blocked           = "blocked"
)

type AuthAuditService struct {
	authAuditRepo interfaces.AuthAuditRepository
}

func NewAuthAuditService(authAuditRepo interfaces.AuthAuditRepository) *AuthAuditService {
	return &AuthAuditService{authAuditRepo: authAuditRepo}
}

func (s AuthAuditService) Create(event string, userID uint) error {
	audit := &models.AuthenticationAudit{
		UserID: userID,
		Event:  event,
		Time:   time.Now(),
	}

	return s.authAuditRepo.Create(audit)
}

func (s AuthAuditService) GetAllAuditsByUserID(userID uint) ([]models.AuthenticationAudit, error) {
	return s.authAuditRepo.GetAllAuditsByUserID(userID)
}
