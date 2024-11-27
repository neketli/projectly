package entity

type Board struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id"`
	Title     string `json:"title"`
}

type BoardTeam struct {
	Board  Board `json:"board"`
	TeamID int   `json:"team_id"`
}
