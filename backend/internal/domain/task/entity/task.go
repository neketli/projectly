package entity

type Task struct {
	ID             int    `json:"id"`
	ProjectIndex   int    `json:"project_index"`
	Title          string `json:"title"`
	Description    string `json:"description,omitempty"`
	Priority       int    `json:"priority,omitempty"`
	StoryPoints    int    `json:"story_points,omitempty"`
	TrackedTime    int    `json:"tracked_time,omitempty"`
	Deadline       int64  `json:"deadline,omitempty"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
	FinishedAt     int64  `json:"finished_at,omitempty"`
	StatusID       int    `json:"status_id"`
	CreatedUserID  int    `json:"created_user_id"`
	AssignedUserID int    `json:"assigned_user_id,omitempty"`
}

type TaskDetailedParams struct {
	TaskID       *int    `json:"task_id" query:"task_id"`
	TeamID       *int    `json:"team_id" query:"team_id"`
	UserID       *int    `json:"user_id" query:"user_id"`
	BoardID      *int    `json:"board_id" query:"board_id"`
	ProjectCode  *string `json:"project_code" query:"project_code"`
	ProjectIndex *int    `json:"project_index" query:"project_index"`
	Limit        *uint64 `json:"limit" query:"limit"`
	Search       *string `json:"search" query:"search"`
}

type TaskDetailed struct {
	ID             int    `json:"id"`
	ProjectCode    string `json:"project_code"`
	ProjectIndex   int    `json:"project_index"`
	Title          string `json:"title"`
	Description    string `json:"description,omitempty"`
	Priority       int    `json:"priority"`
	TrackedTime    int    `json:"tracked_time,omitempty"`
	StoryPoints    int    `json:"story_points"`
	Deadline       int64  `json:"deadline"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
	FinishedAt     int64  `json:"finished_at,omitempty"`
	StatusID       int    `json:"status_id"`
	CreatedUserID  int    `json:"created_user_id"`
	AssignedUserID int    `json:"assigned_user_id"`
	Status         struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		HexColor string `json:"hex_color"`
	} `json:"status"`
	CreatedUser  User `json:"created_user"`
	AssignedUser User `json:"assigned_user,omitempty"`
	Meta         struct {
		TeamID    int `json:"team_id"`
		ProjectID int `json:"project_id"`
		BoardID   int `json:"board_id"`
	} `json:"meta"`
}
