package entity

import (
	"time"
)

type Task struct {
	ID             int        `json:"id"`
	ProjectIndex   int        `json:"project_index"`
	Title          string     `json:"title"`
	Description    *string    `json:"description"`
	Priority       *int       `json:"priority"`
	StoryPoints    *int       `json:"story_points"`
	TrackedTime    *int       `json:"tracked_time"`
	Deadline       *time.Time `json:"deadline"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	FinishedAt     *time.Time `json:"finished_at"`
	StatusID       int        `json:"status_id"`
	CreatedUserID  int        `json:"created_user_id"`
	AssignedUserID *int       `json:"assigned_user_id"`
}

type TaskCard struct {
	ProjectCode  string     `json:"project_code"`
	ProjectIndex int        `json:"project_index"`
	Title        string     `json:"title"`
	Priority     *int       `json:"priority"`
	StoryPoints  *int       `json:"story_points"`
	Deadline     *time.Time `json:"deadline"`
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
