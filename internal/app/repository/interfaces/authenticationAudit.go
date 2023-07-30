package interfaces

import "auth_audit/internal/app/repository/models"

type AuthenticationAuditRepository interface {
	Save(log *models.AuthenticationAudit) error
}
