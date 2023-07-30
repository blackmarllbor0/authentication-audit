package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Login               string `gorm:"not null;unique"`
	Password            string `gorm:"not null"`
	FailedLoginAttempts uint   `gorm:"not null"`
}
