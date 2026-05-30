package entity

const (
	ActionTaskCreated       = "task_created"
	ActionTaskUpdated       = "task_updated"
	ActionStatusChanged     = "status_changed"
	ActionCommentAdded      = "comment_added"
	ActionCommentDeleted    = "comment_deleted"
	ActionAttachmentAdded   = "attachment_added"
	ActionAttachmentDeleted = "attachment_deleted"
)

type TaskActivity struct {
	ID         int    `json:"id"`
	TaskID     int    `json:"task_id"`
	User       User   `json:"user"`
	ActionType string `json:"action_type"`
	FieldName  string `json:"field_name,omitempty"`
	OldValue   string `json:"old_value,omitempty"`
	NewValue   string `json:"new_value,omitempty"`
	CreatedAt  int64  `json:"created_at"`
}
