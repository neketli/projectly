package entity

// Role represents a user role in a team.
type Role struct {
	ID       int    `json:"id"`
	RoleName string `json:"role_name"`
}

// RoleOwner is the owner role with full permissions.
var RoleOwner = &Role{
	ID:       1,
	RoleName: "owner",
}

// RoleEditor is the editor role with edit permissions.
var RoleEditor = &Role{
	ID:       2,
	RoleName: "editor",
}

// RoleDeveloper is the developer role.
var RoleDeveloper = &Role{
	ID:       3,
	RoleName: "developer",
}

// RoleUser is the basic user role.
var RoleUser = &Role{
	ID:       4,
	RoleName: "user",
}