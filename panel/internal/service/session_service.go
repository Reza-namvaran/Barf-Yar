package service

import (
    "crypto/rand"
    "encoding/base64"
    "time"

    "github.com/Reza-namvaran/Barf-Yar/panel/internal/auth"
    "github.com/Reza-namvaran/Barf-Yar/panel/internal/repository"
)

type AuthService interface {
    HashPassword(password string) (string, error)
    CheckPasswordHash(password, hash string) bool
    GenerateSessionToken() (string, error)
    ValidateSessionToken(token string) bool
}

type authService struct {
    repo repository.SessionRepository
}

func NewAuthService(repo repository.SessionRepository) AuthService {
    return &authService{repo: repo}
}

func (s *authService) HashPassword(password string) (string, error) {
    return auth.HashPassword(password)
}

func (s *authService) CheckPasswordHash(password, hash string) bool {
    return auth.CheckPasswordHash(password, hash)
}

func (s *authService) GenerateSessionToken() (string, error) {
    b := make([]byte, 32)
    if _, err := rand.Read(b); err != nil {
        return "", err
    }
    token := base64.URLEncoding.EncodeToString(b)
    expires := time.Now().Add(24 * time.Hour)
    if err := s.repo.Save(token, expires); err != nil {
        return "", err
    }
    return token, nil
}

func (s *authService) ValidateSessionToken(token string) bool {
    expiresAt, err := s.repo.GetExpiry(token)
    if err != nil {
        return false
    }
    return time.Now().Before(expiresAt)
}
