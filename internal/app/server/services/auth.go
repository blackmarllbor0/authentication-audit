package services

import (
	repositories "auth_audit/internal/app/repository/interfaces"
	"auth_audit/internal/app/repository/models"
	services "auth_audit/internal/app/server/services/interfaces"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository repositories.UserRepository
	sessionService services.SessionService
	logAuth        services.LogAuthService
}

func NewAuthService(userRepository repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s AuthService) Auth(login, pwd string) (*models.Session, error) {
	user, err := s.userRepository.FindByLogin(login)
	if err != nil {
		return nil, err
	}

	isComparePwd, err := s.checkPwd(user, pwd)
	if err != nil {
		return nil, err
	}
	if !isComparePwd {
		if err := s.logAuth.LogAuthFailure(user); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("an error occurred during data validation. Make sure. that all information is correct")
	}

	session, err := s.sessionService.CreateSession(user)
	if err != nil {
		return nil, err
	}

	if err := s.logAuth.LogAuthSuccess(user); err != nil {
		return nil, err
	}

	return session, nil
}

func (s AuthService) hashPwd(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(bytes), err
}

func (s AuthService) checkPwd(user *models.User, pwd string) (bool, error) {
	hashPwd, err := s.hashPwd(pwd)
	if err != nil {
		return false, err
	}

	return hashPwd == user.Password, nil
}
