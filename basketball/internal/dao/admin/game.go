package admin

import "basketball/internal/dao"

type GameListReq struct {
	LeagueID uint         `json:"league_id" form:"league_id"`
	PageInfo dao.PageInfo `json:"page_info" form:"page_info"`
}

type GameSubmitReq struct {
	Id          uint   `json:"id" form:"id"`
	LeagueID    uint   `json:"league_id" form:"league_id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	HomeTeamID  uint   `json:"home_team_id" form:"home_team_id"`
	AwayTeamID  uint   `json:"away_team_id" form:"away_team_id"`
	DateTime    string `json:"datetime" form:"datetime"`
	Status      int    `json:"status" form:"status"`
}
