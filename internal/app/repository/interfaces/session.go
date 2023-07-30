package interfaces

import "auth_audit/internal/app/repository/models"

type SessionRepository interface {
	Create(session *models.Session) (id uint, err error)
	FindByToken(token string) (*models.Session, error)
	DeleteByToken(token string) error
}
