package services

import (
	"auth_audit/internal/app/repository/interfaces"
	"auth_audit/internal/app/repository/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s UserService) CreateUser(login, pwd string) (*models.User, error) {
	if len(pwd) < 8 {
		return nil, fmt.Errorf("password must contain at least 8 characters")
	}

	pwdHash, err := s.hashPwd(pwd)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Login:        login,
		PasswordHash: pwdHash,
	}

	if err := s.userRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserService) GetUserByLogin(login string) (*models.User, error) {
	return s.userRepository.GetByLogin(login)
}

func (s UserService) hashPwd(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
