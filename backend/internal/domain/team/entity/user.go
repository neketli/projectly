package entity

import "projectly-server/internal/domain/user/entity"

// TeamUser represents a user within a team with their role.
type TeamUser struct {
	entity.User
	Role Role `json:"role"`
}
