package storage

import (
	"database/sql"
)

// TODO: Move to modele folder
type Activity struct {
	ID        int    `json:"id"`
	MessageID int64  `json:"message_id"`
	Title     string `json:"title"`
}

type activityService struct {
	db *sql.DB
}

type ActivityService interface {
	GetActivityByID(id int) (*Activity, error)
	GetAllActivities() ([]*Activity, error)
	CountActivities() (int, error)
	// AddActivity()
	// DeleteActivity()
}

func NewActivityService(db *sql.DB) ActivityService {
	return &activityService{db: db}
}

func (s *activityService) GetActivityByID(id int) (*Activity, error) {
	activity := &Activity{}
	err := s.db.QueryRow(`
		SELECT id, message_id, title 
		FROM activities 
		WHERE id = $1`, id).Scan(&activity.ID, &activity.MessageID, &activity.Title)
	if err != nil {
		//TODO: cutomize error
		return nil, err
	}

	return activity, nil
}

func (s *activityService) GetAllActivities() ([]*Activity, error) {
	rows, err := s.db.Query(`
		SELECT id, message_id, title
		FROM activities
		ORDER BY id ASC`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var allActivities []*Activity
	for rows.Next() {
		a := &Activity{}
		err = rows.Scan(&a.ID, &a.MessageID, &a.Title)
		if err != nil {
			return nil, err
		}
		allActivities = append(allActivities, a)
	}

	return allActivities, nil
}

func (s *activityService) CountActivities() (int, error) {
	var count int
	err := s.db.QueryRow(`SELECT COUNT(*) FROM activities`).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
