package audits

import (
	"github.com/jinzhu/gorm"
	"time"
)

type AuthenticationAudit struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	EventTime time.Time
	EventType string
}
