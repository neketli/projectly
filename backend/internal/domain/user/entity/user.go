package entity

type User struct {
	ID       int               `json:"id"`
	Name     string            `json:"name"`
	Surname  string            `json:"surname"`
	Email    string            `json:"email"`
	Password string            `json:"-"`
	Meta     map[string]string `json:"-"`
}
