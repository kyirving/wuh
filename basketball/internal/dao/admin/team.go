package admin

// 用于接口的 输入、输出参数
type TeamDao struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LeagueID    int64  `json:"league_id"`
}

type TeamListResp struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LeagueID    int64  `json:"league_id"`
	LeagueName  string `json:"league_name"`
	Status      int64  `json:"status"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}
