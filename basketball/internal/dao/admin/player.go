package admin

type PlayerReq struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TeamID      int    `json:"team_id"`
	Number      int    `json:"number"`
	Position    string `json:"position"`
	Height      int    `json:"height"`
	Weight      int    `json:"weight"`
}

type PlayerListReq struct {
	Name     string `json:"name"`
	PageNum  int    `json:"page_num" default:"1"`
	PageSize int    `json:"page_size" default:"10"`
}

type PlayerListItem struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TeamID      int    `json:"team_id"`
	Number      int    `json:"number"`
	Position    string `json:"position"`
	Height      int    `json:"height"`
	Weight      int    `json:"weight"`
	TeamName    string `json:"team_name"`
	CreatedAt   string `json:"create_time"`
}
type PlayerListResp struct {
	Players []PlayerListItem `json:"players"`
	Total   int64            `json:"total"`
}
