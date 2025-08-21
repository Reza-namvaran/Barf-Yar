package repository

import (
	"database/sql"
	"time"
)

type SessionRepository interface {
	Save(token string, expiresAt time.Time) error
	GetExpiry(token string) (time.Time, error)
	Delete(token string) error
}

type sessionRepo struct {
    db *sql.DB
}

func NewSessionRepository(db *sql.DB) SessionRepository {
    return &sessionRepo{db: db}
}

func (repo *sessionRepo) Save(token string, expiresAt time.Time) error {
    _, err := repo.db.Exec(`INSERT INTO sessions (token, expires_at) VALUES ($1, $2)`, token, expiresAt)
    if err != nil {
		return ErrSaveSession
	}
	
	return nil
}

func (repo *sessionRepo) GetExpiry(token string) (time.Time, error) {
	var expiresAt time.Time
	err := repo.db.QueryRow(`SELECT expires_at FROM sessions WHERE token=$1`, token).Scan(&expiresAt)
	if err != nil {
		return expiresAt, ErrFailedToFetch
	}

	return expiresAt, nil
}

func (repo *sessionRepo) Delete(token string) error {
    _, err := repo.db.Exec(`DELETE FROM sessions WHERE token=$1`, token)
	if err != nil {
		return ErrDeleteSession
	}

    return nil
}
