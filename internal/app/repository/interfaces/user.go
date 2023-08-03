package interfaces

import "auth_audit/internal/app/repository/models"

type UserRepository interface {
	Create(user *models.User) error
	GetByLogin(login string) (*models.User, error)
	GetById(ID uint) (*models.User, error)
	Block()
}
