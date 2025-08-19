package repository

import (
	"database/sql"
	"errors"	

	"github.com/Reza-namvaran/Barf-Yar/panel/internal/models"
)

type ActivityRepository interface {
	GetByID(id int) (*models.Activity, error)
	GetAll() ([]*models.Activity, error)
	Insert(activity *models.Activity) (int, error)
	LinkSupportPrompt(activityID, supportPromptID int) error
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
		SELECT a.id, a.message_id, a.title, ap.prompt_message_id FROM activities a
		LEFT JOIN activity_prompts ap ON a.id = ap.activity_id
		ORDER BY a.id ASC`)
	if err != nil {
		return nil, errors.New("Failed to fetch all activities")
	}
	defer rows.Close()
	
	var allActivities []*models.Activity
	for rows.Next() {
		activity := &models.Activity{}
		err = rows.Scan(&activity.ID, &activity.MessageID, &activity.Title, &activity.PromptMessageID)
		if err != nil {
			return nil, errors.New("Failed to create activity")
		}

		allActivities = append(allActivities, activity)
	}

	return allActivities, nil
}

func (repo *activityRepo) Insert(activity *models.Activity) (int, error) {
	var id int
	err := repo.db.QueryRow(`
		INSERT INTO activities (message_id, title)
		VALUES ($1, $2)
		RETURNING id`, activity.MessageID, activity.Title).Scan(&id)
	if err != nil {
		return 0, errors.New("Could not save activity")
	}

	return id, nil
}

func (repo *activityRepo) LinkSupportPrompt(activityID, supportPromptID int) error {
	_, err := repo.db.Exec(`
		INSERT INTO activity_prompts (activity_id, prompt_message_id)
		VALUES($1, $2)
		ON CONFLICT (activity_id) DO UPDATE
		SET prompt_message_id = EXCLUDED.prompt_message_id`, activityID, supportPromptID)
	
	if err != nil {
		return errors.New("Could not link prompt to activity")
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