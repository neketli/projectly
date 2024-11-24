package entity

type Role struct {
	ID       int    `json:"id"`
	RoleName string `json:"role_name"`
}

var (
	RoleOwner = &Role{
		ID:       1,
		RoleName: "owner",
	}
	RoleEditor = &Role{
		ID:       2,
		RoleName: "editor",
	}
	RoleDeveloper = &Role{
		ID:       3,
		RoleName: "developer",
	}
	RoleUser = &Role{
		ID:       4,
		RoleName: "user",
	}
)
