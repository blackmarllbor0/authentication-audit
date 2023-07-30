package interfaces

import "auth_audit/internal/app/repository/models"

type LogAuthService interface {
	LogAuthFailure(user *models.User) error
	LogAuthSuccess(user *models.User) error
}
