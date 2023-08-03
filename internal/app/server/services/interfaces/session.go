package interfaces

import "auth_audit/internal/app/repository/models"

type SessionService interface {
	Create(userID uint) (*models.Session, error)
}
