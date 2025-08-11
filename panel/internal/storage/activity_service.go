package storage


type Activity struct {
	ID        int    `json:"id"`
	MessageID int64  `json:"message_id"`
	Title     string `json:"title"`
}

type activityService struct {
	db *sql.DB
}