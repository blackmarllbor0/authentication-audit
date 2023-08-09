package interfaces

import (
	"auth_audit/internal/app/repository/models"
	"auth_audit/internal/app/server/DTO"
)

type UserService interface {
	CreateUser(createUserDTO DTO.RegisterUserDTO) (*models.User, error)
	GetUserByLogin(login string) (*models.User, error)
}
