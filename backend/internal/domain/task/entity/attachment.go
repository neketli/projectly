package entity

type Attachment struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	TaskID int    `json:"task_id"`
}
