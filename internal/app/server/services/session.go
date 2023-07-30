package services

import (
	repositories "auth_audit/internal/app/repository/interfaces"
	"auth_audit/internal/app/repository/models"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"strconv"
	"time"
)

const (
	sessionDuration = time.Hour * 1
)

type SessionService struct {
	sessionRepo repositories.SessionRepository
}

func NewSessionService(sessionRepository repositories.SessionRepository) *SessionService {
	return &SessionService{sessionRepo: sessionRepository}
}

func (s SessionService) CreateSession() (*models.Session, error) {
	token, err := s.generateSessionToken()
	if err != nil {
		return nil, err
	}

	session := &models.Session{
		Token:     token,
		ExpiresAt: time.Now().Add(sessionDuration),
	}

	sessionID, err := s.sessionRepo.Create(session)

	session.ID = sessionID

	return session, nil
}

func (s SessionService) generateSessionToken() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	salt := fmt.Sprintf("%x", b)

	timestamp := time.Now().Unix()

	data := salt + strconv.FormatInt(timestamp, 10)

	hasher := sha1.New()
	hasher.Write([]byte(data))
	sha1Hash := hasher.Sum(nil)

	token := fmt.Sprintf("%x", sha1Hash)

	return token, nil
}
