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

type TaskCard struct {
	ProjectCode  string `json:"project_code"`
	ProjectIndex int    `json:"project_index"`
	Title        string `json:"title"`
	Priority     int    `json:"priority"`
	StoryPoints  int    `json:"story_points"`
	Deadline     int    `json:"deadline"`
	Status       struct {
		Title    string `json:"title"`
		HexColor string `json:"hex_color"`
	} `json:"status"`
	AssignedUser struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
		Avatar  string `json:"avatar"`
	} `json:"assigned_user"`
}
