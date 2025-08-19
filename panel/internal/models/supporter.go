package models

import "time"

type Supporter struct {
	ID         int       `json:"id"`
	ActivityID int       `json:"activity_id"`
	UserID     int64     `json:"user_id"`
	JoinedAt   time.Time `json:"joined_at"`
}