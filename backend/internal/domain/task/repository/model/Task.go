package model

import (
	"database/sql"
	"task-tracker-server/internal/domain/task/entity"
)

type Task struct {
	ID             int            `db:"id"`
	ProjectIndex   int            `db:"project_index"`
	Title          string         `db:"title"`
	Description    sql.NullString `db:"description"`
	Priority       sql.NullInt64  `db:"priority"`
	StoryPoints    sql.NullInt64  `db:"story_points"`
	TrackedTime    sql.NullInt64  `db:"tracked_time"`
	Deadline       sql.NullTime   `db:"deadline"`
	CreatedAt      sql.NullTime   `db:"created_at"`
	UpdatedAt      sql.NullTime   `db:"updated_at"`
	FinishedAt     sql.NullTime   `db:"finished_at"`
	StatusID       int            `db:"status_id"`
	CreatedUserID  int            `db:"created_user_id"`
	AssignedUserID sql.NullInt64  `db:"assigned_user_id"`
}

func (t Task) ToEntity() entity.Task {
	var finishedAt int64
	if t.FinishedAt.Valid {
		finishedAt = t.FinishedAt.Time.Unix()
	}

	var deadline int64
	if t.Deadline.Valid {
		deadline = t.Deadline.Time.Unix()
	}

	return entity.Task{
		ID:             t.ID,
		ProjectIndex:   t.ProjectIndex,
		Title:          t.Title,
		Description:    t.Description.String,
		Priority:       int(t.Priority.Int64),
		StoryPoints:    int(t.StoryPoints.Int64),
		TrackedTime:    int(t.TrackedTime.Int64),
		Deadline:       deadline,
		CreatedAt:      t.CreatedAt.Time.Unix(),
		UpdatedAt:      t.UpdatedAt.Time.Unix(),
		FinishedAt:     finishedAt,
		StatusID:       t.StatusID,
		CreatedUserID:  t.CreatedUserID,
		AssignedUserID: int(t.AssignedUserID.Int64),
	}
}
