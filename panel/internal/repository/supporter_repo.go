package repository

import (
	"database/sql"
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/models"
)

type SupporterRepository interface {
	GetByActivityID(activityID int) ([]*models.Supporter, error)
}

type supporterRepository struct {
	db *sql.DB
}

func NewSupporterRepository(db *sql.DB) SupporterRepository {
	return &supporterRepository{db: db}
}

func (r *supporterRepository) GetByActivityID(activityID int) ([]*models.Supporter, error) {
	query := `SELECT id, activity_id, user_id, joined_at 
	          FROM activity_supporters WHERE activity_id = $1 
			  ORDER BY joined_at DESC`
	rows, err := r.db.Query(query, activityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var supporters []*models.Supporter
	for rows.Next() {
		var s models.Supporter
		err := rows.Scan(&s.ID, &s.ActivityID, &s.UserID, &s.JoinedAt)
		if err != nil {
			return nil, err
		}
		supporters = append(supporters, &s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return supporters, nil
}