package services

import (
	"auth_audit/internal/app/repository/interfaces"
	"auth_audit/internal/app/repository/models"
	"auth_audit/internal/app/server/DTO"
	"auth_audit/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s UserService) CreateUser(createUserDTO DTO.RegisterUserDTO) (*models.User, error) {
	if len(createUserDTO.Password) < 8 {
		return nil, errors.ShortPassword
	}

	pwdHash, err := s.hashPwd(createUserDTO.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Login:        createUserDTO.Login,
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

func (s UserService) BlockUser(userID uint) error {
	return s.userRepository.Block(userID)
}

func (s UserService) IncrementFailedLoginAttempts(userID uint) (int, error) {
	return s.userRepository.IncrementFailedLoginAttempts(userID)
}

func (s UserService) hashPwd(pwd string) (string, error) {
	if len(pwd) < 8 {
		return "", errors.ShortPassword
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
