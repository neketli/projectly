package entity

type Project struct {
	ID          int    `json:"id"`
	TeamID      int    `json:"team_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Code        string `json:"code"`
}
