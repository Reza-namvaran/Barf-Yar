package repository

import (
	"database/sql"

	"github.com/Reza-namvaran/Barf-Yar/panel/internal/models"
)

type AdminRepository interface {
	GetByUsername(username string) (*models.Admin, error)
	Create(admin *models.Admin) error
}

type adminRepo struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) AdminRepository {
	return &adminRepo{db: db}
}

func (repo *adminRepo) GetByUsername(username string) (*models.Admin, error) {
	admin := &models.Admin{}
	err := repo.db.QueryRow(`
		SELECT id, username, password_hash 
		FROM admins 
		WHERE username = $1`, username).Scan(&admin.ID, &admin.Username, &admin.PasswordHash)

	if err != nil {
		return nil, ErrFailedToFetch
	}

	return admin, nil
}

func (repo *adminRepo) Create(admin *models.Admin) error {
    _, err := repo.db.Exec(`
        INSERT INTO admins (username, password_hash)
        VALUES ($1, $2)
        ON CONFLICT (username) DO NOTHING`,
        admin.Username, admin.PasswordHash,)
	
	if err != nil {
		return ErrCreateAdmin
	}

	return nil
}