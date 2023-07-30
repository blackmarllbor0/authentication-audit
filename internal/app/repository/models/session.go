package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Session struct {
	gorm.Model

	Token     string
	ExpiresAt time.Time
}
