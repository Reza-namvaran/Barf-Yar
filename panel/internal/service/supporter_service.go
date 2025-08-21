package service

import (
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/models"
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/repository"
)

type SupporterService interface {
	GetSupportersByActivity(activityID int) ([]*models.Supporter, error)
}

type supporterService struct {
	repo repository.SupporterRepository
}

func NewSupporterService(repo repository.SupporterRepository) SupporterService {
	return &supporterService{repo: repo}
}

func (s *supporterService) GetSupportersByActivity(activityID int) ([]*models.Supporter, error) {
	return s.repo.GetByActivityID(activityID)
}