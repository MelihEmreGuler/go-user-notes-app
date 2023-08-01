package models

type Note struct {
	ID        string `json:"note_id"`
	UserId    string `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
