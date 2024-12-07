package entity

import "projectly-server/internal/domain/user/entity"

type TeamUser struct {
	entity.User
	Role Role `json:"role"`
}
