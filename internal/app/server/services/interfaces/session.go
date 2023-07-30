package interfaces

import "auth_audit/internal/app/repository/models"

type SessionService interface {
	CreateSession() (*models.Session, error)
}
