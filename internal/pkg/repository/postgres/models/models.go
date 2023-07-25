package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type (
	User struct {
		gorm.Model

		Login               string `gorm:"not null;unique" json:"login"`
		Password            string `gorm:"not null" json:"password"`
		FailedLoginAttempts int    `json:"failedLoginAttempts"`

		Sessions             []Session             `json:"sessions"`
		AuthenticationAudits []AuthenticationAudit `json:"authenticationAudits"`
	}
	Session struct {
		gorm.Model

		UserID         uint      `json:"userID"`
		Token          string    `gorm:"not null" json:"token"`
		ExpirationTime time.Time `gorm:"not null" json:"expirationTime"`

		User User `gorm:"foreignKey:UserID"`
	}
	AuthenticationAudit struct {
		gorm.Model

		UserID    uint      `json:"userID"`
		EventTime time.Time `gorm:"not null" json:"eventTime"`
		EventType string    `gorm:"not null" json:"eventType"`

		User User `gorm:"foreignKey:UserID"`
	}
)
