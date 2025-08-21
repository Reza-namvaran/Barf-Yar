package service

import (
	"database/sql"
	"errors"

	"github.com/Reza-namvaran/Barf-Yar/panel/internal/auth"
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/models"
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/repository"
)


type AdminService interface {
	ValidateAdmin(username, password string) (bool, error)
	GetAdminByUsername(username string) (*models.Admin, error)
	CreateAdmin(username, password string) error
}

type adminService struct {
	repo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
    return &adminService{repo: repo}
}

func (serv *adminService) ValidateAdmin(username, password string) (bool, error) {
	admin, err := serv.GetAdminByUsername(username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return auth.CheckPasswordHash(password, admin.PasswordHash), nil
}

func (serv *adminService) GetAdminByUsername(username string) (*models.Admin, error) {
	return serv.repo.GetByUsername(username)
}

func (serv *adminService) CreateAdmin(username, password string) error {
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return errors.New("Failed to hash provided password")
	}

	admin := &models.Admin{
		Username: username,
		PasswordHash: hashedPassword,
	}
	
	return serv.repo.Create(admin)
}