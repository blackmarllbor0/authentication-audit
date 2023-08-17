package postgres

import (
	"auth_audit/internal/app/repository/models"
	"github.com/jinzhu/gorm"
)

type AuthenticationAudit struct {
	db *gorm.DB
}

func NewAuthenticationAudit(db *gorm.DB) *AuthenticationAudit {
	return &AuthenticationAudit{db: db}
}

func (r AuthenticationAudit) Create(audit *models.AuthenticationAudit) error {
	return r.db.Create(&audit).Error
}

func (r AuthenticationAudit) GetAllAuditsByUserID(userID uint) ([]models.AuthenticationAudit, error) {
	var auditEntries []models.AuthenticationAudit

	if err := r.db.Where("user_id = ?", userID).Find(&auditEntries).Error; err != nil {
		return nil, err
	}

	return auditEntries, nil
}

func (r AuthenticationAudit) ClearAuthAuditsByToken(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.AuthenticationAudit{}).Error
}
