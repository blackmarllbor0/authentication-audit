package interfaces

import "auth_audit/internal/app/repository/models"

type AuthService interface {
	Auth(login, pwd string) (*models.Session, error)
}
