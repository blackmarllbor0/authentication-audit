package services

import (
	"auth_audit/internal/app/repository/models"
	"auth_audit/internal/app/server/DTO"
	"auth_audit/internal/app/server/services/interfaces"
	"auth_audit/pkg/errors"
	"github.com/jinzhu/gorm"
)

type AuthService struct {
	userService    interfaces.UserService
	sessionService interfaces.SessionService
}

func NewAuthService(userService interfaces.UserService, sessionService interfaces.SessionService) *AuthService {
	return &AuthService{
		userService:    userService,
		sessionService: sessionService,
	}
}

func (s AuthService) Register(dto DTO.RegisterUserDTO) (*models.Session, error) {
	existingUser, err := s.userService.GetUserByLogin(dto.Login)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.UserAlreadyExist
	}

	user, err := s.userService.CreateUser(dto)
	if err != nil {
		return nil, err
	}

	session, err := s.sessionService.Create(user.ID)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s AuthService) Login(login, pwd string) (*models.Session, error) {
	//TODO implement me
	panic("implement me")
}
