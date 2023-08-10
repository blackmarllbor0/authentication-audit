package services

import (
	"auth_audit/internal/app/repository/models"
	"auth_audit/internal/app/server/DTO"
	"auth_audit/internal/app/server/services/interfaces"
	"auth_audit/pkg/errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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

	if dto.Login == "" || dto.Password == "" {
		return nil, errors.MustBeProvidedLoginAndPwd
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

func (s AuthService) Login(dto DTO.LoginUserDTO) (*models.Session, error) {
	if dto.Login == "" || dto.Password == "" {
		return nil, errors.MustBeProvidedLoginAndPwd
	}

	user, err := s.userService.GetUserByLogin(dto.Login)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.InvalidLoginOrPassword
		}

		return nil, err
	}

	if !s.checkPassword(dto.Password, user.PasswordHash) {
		return nil, errors.InvalidLoginOrPassword
	}

	session, err := s.sessionService.Create(user.ID)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s AuthService) checkPassword(pwd, hashPwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd)) == nil
}
