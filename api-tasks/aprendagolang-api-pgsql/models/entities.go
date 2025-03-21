package models

type TaskID int64

type Task struct {
	ID          TaskID `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
