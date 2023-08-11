package interfaces

import "auth_audit/internal/app/repository/models"

type AuthAuditService interface {
	Create(event string, userID uint) error
	GetAllAuditsByUserID(userID uint) ([]models.AuthenticationAudit, error)
}
