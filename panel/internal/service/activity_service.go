package service

import (
    "errors"
    "github.com/Reza-namvaran/Barf-Yar/panel/internal/models"
    "github.com/Reza-namvaran/Barf-Yar/panel/internal/repository"
)

type ActivityService interface {
	GetActivityByID(id int) (*models.Activity, error)
	GetAllActivities() ([]*models.Activity, error)
	CountActivities() (int, error)
	AddActivity(activity *models.Activity) error
	DeleteActivity(id int) error
}

type activityService struct {
	repo repository.ActivityRepository
}

func NewActivityService(repo repository.ActivityRepository) ActivityService {
    return &activityService{repo: repo}
}

func (serv *activityService) GetActivityByID(id int) (*models.Activity, error) {
	return serv.repo.GetByID(id) 
}

func (serv *activityService) GetAllActivities() ([]*models.Activity, error) {
	return serv.repo.GetAll()
}

func (serv *activityService) CountActivities() (int, error) {
	return serv.repo.Count()
}

func (serv *activityService) AddActivity(activity *models.Activity) error {
	// Business rule: no duplicate message_id
	exists, _ := serv.repo.ExistsByID(activity.ID)

	if exists {
		return errors.New("This activity already exists")
	}

	serv.repo.Insert(activity)
	return nil
}

func (serv *activityService) DeleteActivity(id int) error {
	if err := serv.repo.Delete(id); err != nil {
		return errors.New("Can't delete activity")
	}
	
	return nil
}
