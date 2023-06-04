package users

import (
	"github.com/jinzhu/gorm"
	"time"
)

type (
	Users struct {
		gorm.Model
		ID                  uint   `gorm:"primaryKey"`
		Login               string `gorm:"unique"`
		Password            string
		FailedLoginAttempts int
		CreatedAt           time.Time
		UpdatedAt           time.Time
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
		Login:     login,
		Password:  hashPwd,
		CreatedAt: time.Now(),
	}

	if err := r.db.Create(&user).Error; err != nil {
		return Users{}, err
	}

	return user, nil
}
