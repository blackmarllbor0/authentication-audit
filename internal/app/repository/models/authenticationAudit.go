package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type AuthenticationAudit struct {
	gorm.Model
	UserID uint      `gorm:"TYPE:integer REFERENCES users;not null"`
	Time   time.Time `gorm:"not null"`
	Event  string    `gorm:"type:varchar(100);not null"`
}
