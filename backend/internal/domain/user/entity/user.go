package entity

type UserMeta struct {
	Avatar     string `json:"avatar" example:"avatar.png"`
	Provider   string `json:"provider" example:"google"`
	ProviderID string `json:"provider_id" example:"google_id"`
	Language   string `json:"language,omitempty" example:"en"`
	Birthday   string `json:"birthday,omitempty" example:"01.01.2000"`
	Location   string `json:"location,omitempty" example:"Moscow, Russia"`
	About      string `json:"about,omitempty" example:"Something about"`
}

type User struct {
	ID       int       `json:"id" example:"1"`
	Name     string    `json:"name" example:"John"`
	Surname  string    `json:"surname" example:"Doe"`
	Email    string    `json:"email" example:"john.doe@example.com"`
	Password string    `json:"-" example:"password123"`
	Meta     *UserMeta `json:"meta"`
}
