package models

type Activity struct {
	ID        int    `json:"id"`
	MessageID int  `json:"message_id"`
	Title     string `json:"title"`
	PromptMessageID *int `json:"prompt_message_id"`
}
