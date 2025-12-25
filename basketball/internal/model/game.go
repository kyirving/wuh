package model

import "time"

type Game struct {
	Id          int64     `json:"id" gorm:"column:id;primaryKey"`
	Name        string    `json:"name" gorm:"column:name"`
	DateTime    time.Time `json:"date_time" gorm:"column:datetime"`
	Status      int64     `json:"status" gorm:"column:status"`
	Description string    `json:"description" gorm:"column:description"`
	HomeTeamID  uint      `json:"home_team_id" gorm:"column:home_team_id"`
	AwayTeamID  uint      `json:"away_team_id" gorm:"column:away_team_id"`
	HomeTeam    Team      `json:"home_team" gorm:"foreignKey:HomeTeamID;references:Id"`
	AwayTeam    Team      `json:"away_team" gorm:"foreignKey:AwayTeamID;references:Id"`
	LeagueID    uint      `json:"league_id" gorm:"column:league_id"`
	League      League    `json:"league" gorm:"foreignKey:LeagueID;references:Id"`
	BaseModel
}

func (Game) TableName() string {
	return "games"
}
