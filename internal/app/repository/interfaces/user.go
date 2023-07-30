package interfaces

import "auth_audit/internal/app/repository/models"

type UserRepository interface {
	Create(user *models.User) error
	FindByLogin(login string) (*models.User, error)
	IncrementFailedAttempts(user *models.User) error
	ResetFailedAttempts(user *models.User) error
	Block(user *models.User) error
}
