package sessions

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Session struct {
	gorm.Model
	UserID         uint      `gorm:"not null" json:"userID,omitempty"`
	Token          string    `gorm:"not null" json:"token,omitempty"`
	ExpirationTime time.Time `gorm:"not null" json:"expirationTime"`
}
