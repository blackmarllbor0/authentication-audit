package sessions

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Session struct {
	gorm.Model
	ID             uint `gorm:"primaryKey"`
	UserID         uint
	Token          string
	ExpirationTime time.Time
}
