package users

import (
	"github.com/jinzhu/gorm"
)

type (
	Users struct {
		gorm.Model
		Login               string // `gorm:"unique,not null" json:"login,omitempty"`
		Password            string // `gorm:"not null" json:"password,omitempty"`
		FailedLoginAttempts int    // `json:"failedLoginAttempts,omitempty"`
	}
	User struct {
		db *gorm.DB
	}
	Repository interface {
		CreateUser(login, hashPwd string) (Users, error)
	}
)

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (r User) CreateUser(login, hashPwd string) (Users, error) {
	user := Users{
		Login:    login,
		Password: hashPwd,
	}

	if err := r.db.Create(&user).Error; err != nil {
		return Users{}, err
	}

	return user, nil
}
