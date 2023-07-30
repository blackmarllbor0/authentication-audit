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

func (r SessionRepository) Create(session *models.Session) (id uint, err error) {
	if err := r.db.Create(&session).Error; err != nil {
		return 0, err
	}

	return session.ID, nil
}

func (r SessionRepository) FindByToken(token string) (*models.Session, error) {
	var session models.Session
	if err := r.db.Where("token = ?", token).First(&session).Error; err != nil {
		return nil, err
	}

	return &session, nil
}

func (r SessionRepository) DeleteByToken(token string) error {
	return r.db.Where("token = ?", token).Delete(&models.Session{}).Error
}
