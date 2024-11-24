package entity

type Team struct {
	ID          int    `json:"id" example:"1"`
	Name        string `json:"name" example:"example team"`
	Description string `json:"description" example:"My team 1"`
}
