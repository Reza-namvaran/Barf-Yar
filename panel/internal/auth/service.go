package auth

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

type Service interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
	GenerateSessionToken() (string, error)
	ValidateSessionToken(token string) bool
}

type service struct {
	sessions map[string]time.Time
}

func NewService() Service {
	return &service{
		sessions: make(map[string]time.Time),
	}
}

func (s *service) HashPassword(password string) (string, error) {
	return HashPassword(password)
}

func (s *service) CheckPasswordHash(password, hash string) bool {
	return CheckPasswordHash(password, hash)
}

func (s *service) GenerateSessionToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(b)
	s.sessions[token] = time.Now().Add(24 * time.Hour) // 24 hour expiry
	return token, nil
}

func (s *service) ValidateSessionToken(token string) bool {
	if expiry, exists := s.sessions[token]; exists {
		if time.Now().Before(expiry) {
			return true
		}
		delete(s.sessions, token)
	}
	return false
}
