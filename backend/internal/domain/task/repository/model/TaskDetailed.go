package model

import (
	"database/sql"
	"projectly-server/internal/domain/task/entity"
)

type TaskDetailed struct {
	ID             int           `json:"id"`
	ProjectCode    string        `json:"project_code"`
	ProjectIndex   int           `json:"project_index"`
	Title          string        `json:"title"`
	Description    string        `json:"description,omitempty"`
	Priority       sql.NullInt64 `json:"priority"`
	TrackedTime    sql.NullInt64 `json:"tracked_time,omitempty"`
	StoryPoints    sql.NullInt64 `json:"story_points"`
	Deadline       sql.NullTime  `json:"deadline"`
	CreatedAt      sql.NullTime  `json:"created_at"`
	UpdatedAt      sql.NullTime  `json:"updated_at"`
	FinishedAt     sql.NullTime  `json:"finished_at,omitempty"`
	StatusID       int           `json:"status_id"`
	CreatedUserID  int           `json:"created_user_id"`
	AssignedUserID sql.NullInt64 `json:"assigned_user_id,omitempty"`
	Status         struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		HexColor string `json:"hex_color"`
	} `json:"status"`
	CreatedUser struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Surname string `json:"surname"`
		Email   string `json:"email"`
		Avatar  string `json:"avatar"`
	} `json:"created_user"`
	AssignedUser struct {
		ID      sql.NullInt64  `json:"id,omitempty"`
		Name    sql.NullString `json:"name,omitempty"`
		Surname sql.NullString `json:"surname,omitempty"`
		Email   sql.NullString `json:"email,omitempty"`
		Avatar  string         `json:"avatar,omitempty"`
	} `json:"assigned_user,omitempty"`
	Meta struct {
		TeamID    int `json:"team_id"`
		ProjectID int `json:"project_id"`
		BoardID   int `json:"board_id"`
	} `json:"meta"`
}

func (t TaskDetailed) ToEntity() entity.TaskDetailed {
	task := entity.TaskDetailed{
		ID:             t.ID,
		ProjectCode:    t.ProjectCode,
		ProjectIndex:   t.ProjectIndex,
		Title:          t.Title,
		Description:    t.Description,
		Priority:       int(t.Priority.Int64),
		TrackedTime:    int(t.TrackedTime.Int64),
		StoryPoints:    int(t.StoryPoints.Int64),
		Deadline:       0,
		CreatedAt:      t.CreatedAt.Time.Unix(),
		UpdatedAt:      t.UpdatedAt.Time.Unix(),
		FinishedAt:     0,
		StatusID:       t.StatusID,
		CreatedUserID:  t.CreatedUserID,
		AssignedUserID: 0,
	}

	if t.FinishedAt.Valid {
		task.FinishedAt = t.FinishedAt.Time.Unix()
	}

	if t.Deadline.Valid {
		task.Deadline = t.Deadline.Time.Unix()
	}

	if t.AssignedUserID.Valid {
		task.AssignedUserID = int(t.AssignedUserID.Int64)
	}

	task.Status.ID = t.Status.ID
	task.Status.Title = t.Status.Title
	task.Status.HexColor = t.Status.HexColor

	task.CreatedUser.ID = t.CreatedUser.ID
	task.CreatedUser.Name = t.CreatedUser.Name
	task.CreatedUser.Surname = t.CreatedUser.Surname
	task.CreatedUser.Email = t.CreatedUser.Email
	task.CreatedUser.Avatar = t.CreatedUser.Avatar

	task.AssignedUser.ID = int(t.AssignedUser.ID.Int64)
	task.AssignedUser.Name = t.AssignedUser.Name.String
	task.AssignedUser.Surname = t.AssignedUser.Surname.String
	task.AssignedUser.Email = t.AssignedUser.Email.String
	task.AssignedUser.Avatar = t.AssignedUser.Avatar

	task.Meta.TeamID = t.Meta.TeamID
	task.Meta.ProjectID = t.Meta.ProjectID
	task.Meta.BoardID = t.Meta.BoardID

	return task
}
