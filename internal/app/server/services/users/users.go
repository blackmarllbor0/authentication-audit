package users

import (
	"auth_audit/internal/pkg/repository/postgres/models"
	"auth_audit/internal/pkg/repository/postgres/models/users"
	"crypto/sha256"
	"fmt"
)

type (
	Users struct {
		userRepo users.Repository
	}
	Service interface {
		CreateUser(login, pwd string) (models.User, error)
	}
)

func NewService(repo users.Repository) *Users {
	return &Users{userRepo: repo}
}

func (s Users) CreateUser(login, pwd string) (models.User, error) {
	if len(pwd) < 8 {
		return models.User{}, fmt.Errorf("the number of characters must be less than 8")
	}

	return s.userRepo.CreateUser(login, s.hashPwd(pwd))
}

func (s Users) hashPwd(pwd string) string {
	hasher := sha256.New()
	hasher.Write([]byte(pwd))
	hashPwd := hasher.Sum(nil)

	// Convert hash to hex string.
	hashPwdStr := fmt.Sprintf("%x", hashPwd)
	return hashPwdStr
}

func (s Users) checkHashPwd(pwd, hashPwd string) bool {
	newHash := s.hashPwd(pwd)

	return newHash == hashPwd
}
