package admin

type PlayerReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	TeamID      uint   `json:"team_id"`
	Number      uint   `json:"number"`
	Position    string `json:"position"`
	Height      uint   `json:"height"`
	Weight      uint   `json:"weight"`
}
