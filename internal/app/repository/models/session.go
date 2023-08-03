package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Session struct {
	gorm.Model
	UserID   uint      `gorm:"TYPE:integer REFERENCES users;not null"`
	Token    string    `gorm:"uniqueIndex;not null"`
	LiveTime time.Time `gorm:"not null"`
}
