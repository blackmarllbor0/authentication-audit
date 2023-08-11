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

func (r SessionRepository) GetByToken(token string) (*models.Session, error) {
	var session models.Session
	if err := r.db.Where("token = ?", token).First(&session).Error; err != nil {
		return nil, err
	}

	return &session, nil
}
