package model

type Team struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	LeagueID    int64  `json:"league_id"`
	Status      int64  `json:"status"`
	BaseModel
}

func (Team) TableName() string {
	return "teams"
}
