package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Login              string `gorm:"type:varchar(100);uniqueIndex;not null"`
	PasswordHash       string `gorm:"type:varchar(100);not null"`
	LoginAttempts      int    `gorm:"default:0"`
	Blocked            bool   `gorm:"default:false"`
	LastLoginAttempt   time.Time
	SessionToken       string `gorm:"type:varchar(100)"`
	SessionTokenExpiry time.Time
}
