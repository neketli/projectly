package entity

type Comment struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	TaskID    int    `json:"task_id"`
	User      User   `json:"user"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
