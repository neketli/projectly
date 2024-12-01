package entity

type Status struct {
	ID       int    `json:"id"`
	BoardID  int    `json:"board_id"`
	Title    string `json:"title"`
	Order    int    `json:"order"`
	HexColor string `json:"hex_color"`
}
