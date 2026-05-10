package entity

// Attachment represents a task attachment entity.
type Attachment struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	TaskID int    `json:"task_id"`
}
