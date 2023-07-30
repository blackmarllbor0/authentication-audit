package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type AuthenticationAudit struct {
	gorm.Model

	UserID uint
	Time   time.Time
	Event  string //  "Successful Login", "Invalid Password", "Account Blocked"
}
