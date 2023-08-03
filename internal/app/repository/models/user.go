package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Login               string `gorm:"uniqueIndex;not null"`
	PasswordHash        string `gorm:"not null"`
	FailedLoginAttempts int    `gorm:"default:0"`
	Blocked             bool   `gorm:"default:false"`
}
