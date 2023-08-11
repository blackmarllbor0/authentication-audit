package interfaces

import "auth_audit/internal/app/repository/models"

type AuthAuditRepository interface {
	Create(audit *models.AuthenticationAudit) error
	GetAllAuditsByUserID(userID uint) ([]*models.AuthenticationAudit, error)
}
