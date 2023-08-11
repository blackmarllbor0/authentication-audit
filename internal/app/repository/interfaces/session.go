package interfaces

import "auth_audit/internal/app/repository/models"

type SessionRepository interface {
	Create(session *models.Session) error
	GetByToken(token string) (*models.Session, error)
}
