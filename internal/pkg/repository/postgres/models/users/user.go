package users

import (
	"auth_audit/internal/pkg/repository/postgres/models"
	"github.com/jinzhu/gorm"
)

type (
	User struct {
		db *gorm.DB
	}
	Repository interface {
		CreateUser(login, hashPwd string) (models.User, error)
	}
)

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (r User) CreateUser(login, hashPwd string) (models.User, error) {
	user := models.User{
		Login:    login,
		Password: hashPwd,
	}

	if err := r.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
