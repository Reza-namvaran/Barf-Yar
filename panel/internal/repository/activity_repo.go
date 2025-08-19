package repository

import (
	"database/sql"
	"errors"	

	"github.com/Reza-namvaran/Barf-Yar/panel/internal/models"
)

type ActivityRepository interface {
	GetByID(id int) (*models.Activity, error)
	GetAll() ([]*models.Activity, error)
	Insert(activity *models.Activity) error
	Delete(id int) error
	Count() (int, error)
	ExistsByMessageID(id int) (bool, error)
}

type activityRepo struct {
	db *sql.DB
}

func NewActivityRepository(db *sql.DB) ActivityRepository {
	return &activityRepo{db: db}
}

func (repo *activityRepo) GetByID(id int) (*models.Activity, error) {
	activity := &models.Activity{}
	err := repo.db.QueryRow(`
		SELECT id, message_id, title FROM activities
		WHERE id = $1`, id).Scan(activity)

	if err != nil {
		return nil, errors.New("Failed to fetch activity")
	}
	return activity, err
}

func (repo *activityRepo) GetAll() ([]*models.Activity, error) {
	rows, err := repo.db.Query(`
	SELECT id, message_id, title FROM activities
	ORDER BY id ASC`)

	if err != nil {
		return nil, errors.New("Failed to fetch all activities")
	}

	defer rows.Close()
	var allActivities []*models.Activity
	for rows.Next() {
		activity := &models.Activity{}
		err = rows.Scan(&activity.ID, &activity.MessageID, &activity.Title)
		if err != nil {
			return nil, errors.New("Failed to create activity")
		}

		allActivities = append(allActivities, activity)
	}

	return allActivities, nil
}

func (repo *activityRepo) Insert(activity *models.Activity) error {
	_, err := repo.db.Exec(`
	INSERT INTO activities (message_id, title)
	VALUES ($1, $2)`, activity.MessageID, activity.Title)
	if err != nil {
		return errors.New("Could not save activity")
	}

	return nil
}

func (repo *activityRepo) Delete(id int) error {
	_, err := repo.db.Exec(`DELETE FROM activities WHERE id = $1`, id)
	return err
}

func (repo *activityRepo) Count() (int, error) {
	var count int
	err := repo.db.QueryRow(`SELECT COUNT(*) FROM activities`).Scan(&count)
	if err != nil {
		return 0, errors.New("Can't fecth data")
	}

	return count, nil
}

func (repo *activityRepo) ExistsByMessageID(id int) (bool, error) {
	var exists bool
	err := repo.db.QueryRow(`
	SELECT EXISTS(SELECT 1 FROM activities WHERE message_id = $1)`,
	id,).Scan(&exists)

	if err != nil {
		return false, errors.New("Can't fetch data")
	}

	return exists, nil
}