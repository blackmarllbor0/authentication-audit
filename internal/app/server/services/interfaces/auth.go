package interfaces

import (
	"auth_audit/internal/app/repository/models"
	"auth_audit/internal/app/server/DTO"
)

type AuthService interface {
	Register(dto DTO.RegisterUserDTO) (*models.Session, error)
	Login(dto DTO.LoginUserDTO) (*models.Session, error)
	GetAuthAuditByToken(token string) ([]DTO.AuthAuditDTO, error)
}
