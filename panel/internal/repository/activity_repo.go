package repository

import (
	"database/sql"

	"github.com/Reza-namvaran/Barf-Yar/panel/internal/models"
)

type ActivityRepository interface {
	GetByID(id int) (*models.Activity, error)
	GetAll() ([]*models.Activity, error)
	Insert(activity *models.Activity) (int, error)
	LinkSupportPrompt(activityID, supportPromptID int) error
	RemoveSupportPrompt(activityID int) error
	Update(activity *models.Activity) error 
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
    rows, err := repo.db.Query(`
        SELECT a.id, a.message_id, a.title, ap.prompt_message_id FROM activities a
        LEFT JOIN activity_prompts ap ON a.id = ap.activity_id
        WHERE a.id = $1`, id)

    if err != nil {
        return nil, ErrFetchActivity
    }
    defer rows.Close()

    activity := &models.Activity{}
    found := false

    for rows.Next() {
        found = true
        err := rows.Scan(
            &activity.ID,
            &activity.MessageID,
            &activity.Title,
            &activity.PromptMessageID,
        )
        if err != nil {
            return nil, ErrFailedToScan
        }
        break
    }

    if err = rows.Err(); err != nil {
        return nil, ErrIteration
    }

    if !found {
        return nil, nil
    }

    return activity, nil
}

func (repo *activityRepo) GetAll() ([]*models.Activity, error) {
	rows, err := repo.db.Query(`
		SELECT a.id, a.message_id, a.title, ap.prompt_message_id FROM activities a
		LEFT JOIN activity_prompts ap ON a.id = ap.activity_id
		ORDER BY a.id ASC`)
	if err != nil {
		return nil, ErrAllActivities
	}
	defer rows.Close()
	
	var allActivities []*models.Activity
	for rows.Next() {
		activity := &models.Activity{}
		err = rows.Scan(&activity.ID, &activity.MessageID, &activity.Title, &activity.PromptMessageID)
		if err != nil {
			return nil, ErrCreateActivity
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
		return 0, ErrSaveActivity
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
		return ErrLinkActivity
	}

	return nil
}

func (repo *activityRepo) RemoveSupportPrompt(activityID int) error {
	_, err := repo.db.Exec(`DELETE FROM activity_prompts WHERE activity_id = $1`, activityID)
	if err != nil {
		return ErrRemoveLink
	}

	return nil
}

func (repo *activityRepo) Update(activity *models.Activity) error {
	_, err := repo.db.Exec(`
		UPDATE activities
        SET message_id = $1, title = $2
        WHERE id = $3
		`, activity.MessageID, activity.Title, activity.ID)
	
	if err != nil {
		return ErrUpdateActivity
	}

	if activity.PromptMessageID != nil {
        return repo.LinkSupportPrompt(activity.ID, *activity.PromptMessageID)
    } else {
		err = repo.RemoveSupportPrompt(activity.ID)
		if err != nil {
			return err
		}
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
		return 0, ErrFailedToFetch
	}

	return count, nil
}

func (repo *activityRepo) ExistsByMessageID(id int) (bool, error) {
	var exists bool
	err := repo.db.QueryRow(`
		SELECT EXISTS(SELECT 1 FROM activities WHERE message_id = $1)`,
	id,).Scan(&exists)

	if err != nil {
		return false, ErrFailedToFetch
	}

	return exists, nil
}