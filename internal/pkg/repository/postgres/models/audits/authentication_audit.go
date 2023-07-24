package audits

import (
	"github.com/jinzhu/gorm"
	"time"
)

type AuthenticationAudit struct {
	gorm.Model
	EventTime time.Time `gorm:"not null" json:"eventTime"`
	EventType string    `gorm:"not null" json:"eventType,omitempty"`
}
