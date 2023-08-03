package interfaces

import "auth_audit/internal/app/repository/models"

type UserService interface {
	CreateUser(login, pwd string) (*models.User, error)
	GetUserByLogin(login string) (*models.User, error)
}
