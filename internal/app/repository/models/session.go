package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Session struct {
	gorm.Model

	UserID uint      `gorm:"TYPE:integer REFERENCES users;not null"`
	Token  string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Expiry time.Time `gorm:"not null"`
}
