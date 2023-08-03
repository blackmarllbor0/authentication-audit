package postgres

import (
	"auth_audit/internal/app/repository/models"
	"github.com/jinzhu/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r SessionRepository) Create(session *models.Session) error {
	return r.db.Create(&session).Error
}
