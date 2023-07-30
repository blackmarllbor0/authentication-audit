package postgres

import (
	"auth_audit/internal/app/repository/models"
	"github.com/jinzhu/gorm"
)

func NewAuthenticationAudit(db *gorm.DB) *AuthenticationAudit {
	return &AuthenticationAudit{db: db}
}

type AuthenticationAudit struct {
	db *gorm.DB
}

func (r AuthenticationAudit) Save(log *models.AuthenticationAudit) error {
	return r.db.Create(&log).Error
}
