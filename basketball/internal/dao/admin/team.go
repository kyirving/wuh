package admin

type TeamDao struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	LeagueID    int64  `json:"league_id"`
}
