package services

import (
	"auth_audit/internal/app/repository/interfaces"
	"auth_audit/internal/app/repository/models"
	"auth_audit/pkg/errors"
	"crypto/rand"
	"encoding/base64"
	"time"
)

type SessionService struct {
	sessionRepository interfaces.SessionRepository
}

func NewSessionService(sessionRepository interfaces.SessionRepository) *SessionService {
	return &SessionService{sessionRepository: sessionRepository}
}

func (s SessionService) Create(userID uint) (*models.Session, error) {
	if userID == 0 {
		return nil, errors.NullForeignKey
	}

	token := s.generateToken()
	for len(token) == 0 {
		token = s.generateToken()
	}

	session := &models.Session{
		Token:    token,
		LiveTime: time.Now().Add(time.Hour * 1),
		UserID:   userID,
	}

	if err := s.sessionRepository.Create(session); err != nil {
		return nil, err
	}

	return session, nil
}

func (s SessionService) GetByToken(token string) (*models.Session, error) {
	if len(token) == 0 {
		return nil, errors.TokenIsEmpty
	}

	session, err := s.sessionRepository.GetByToken(token)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s SessionService) ValidateToken(token string) error {
	if token == "" {
		return errors.TokenIsEmpty
	}

	session, err := s.GetByToken(token)
	if err != nil {
		return err
	}

	if !session.LiveTime.After(time.Now()) {
		return errors.TokenHasExpired
	}

	return nil
}

func (s SessionService) generateToken() string {
	buffer := make([]byte, 32)
	_, err := rand.Read(buffer)
	if err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(buffer)[:32]
}
