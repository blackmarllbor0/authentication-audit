package users

import (
	"aptekaaprel/internal/pkg/repository/postgres/models/users"
	"crypto/sha256"
	"fmt"
)

type (
	Users struct {
		userRepo users.Repository
	}
	Service interface {
		CreateUser(login, pwd string) (users.Users, error)
	}
)

func NewService(repo users.Repository) *Users {
	return &Users{userRepo: repo}
}

func (s Users) CreateUser(login, pwd string) (users.Users, error) {
	if len(pwd) < 8 {
		return users.Users{}, fmt.Errorf("the number of characters must be less than 8")
	}

	return s.userRepo.CreateUser(login, s.hashPwd(pwd))
}

func (s Users) hashPwd(pwd string) string {
	hasher := sha256.New()
	hasher.Write([]byte(pwd))
	hashPwd := hasher.Sum(nil)

	// Преоброзование хеша в строку 16-ричного формата.
	hashPwdStr := fmt.Sprintf("%x", hashPwd)
	return hashPwdStr
}

func (s Users) checkHashPwd(pwd, hashPwd string) bool {
	newHash := s.hashPwd(pwd)

	return newHash == hashPwd
}
