package storage

import (
	"database/sql"
	"errors"

	"github.com/Reza-namvaran/Barf-Yar/panel/internal/auth"
)

type AdminService interface {
	ValidateAdmin(username, password string) (bool, error)
	GetAdminByUsername(username string) (*Admin, error)
}

type Admin struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
}

type adminService struct {
	db *sql.DB
}

func NewAdminService(db *sql.DB) AdminService {
	return &adminService{db: db}
}


func (s *adminService) ValidateAdmin(username, password string) (bool, error) {
	admin, err := s.GetAdminByUsername(username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return auth.CheckPasswordHash(password, admin.PasswordHash), nil
}

func (s *adminService) GetAdminByUsername(username string) (*Admin, error) {
	admin := &Admin{}
	err := s.db.QueryRow(`
		SELECT id, username, password_hash 
		FROM admins 
		WHERE username = $1`, username).Scan(&admin.ID, &admin.Username, &admin.PasswordHash)

	if err != nil {
		return nil, err
	}

	return admin, nil
}
