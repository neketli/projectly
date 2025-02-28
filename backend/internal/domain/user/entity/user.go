package entity

type UserMeta struct {
	Avatar     string `json:"avatar" example:"avatar.png"`
	Provider   string `json:"provider" example:"google"`
	ProviderID string `json:"provider_id" example:"google_id"`
}

type User struct {
	ID       int       `json:"id" example:"1"`
	Name     string    `json:"name" example:"John"`
	Surname  string    `json:"surname" example:"Doe"`
	Email    string    `json:"email" example:"john.doe@example.com"`
	Password string    `json:"-" example:"password123"`
	Meta     *UserMeta `json:"meta"`
}
