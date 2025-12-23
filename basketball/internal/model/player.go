package model

type Player struct {
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	TeamID      uint   `json:"team_id" gorm:"column:team_id"`
	Number      uint   `json:"number" gorm:"column:number"`
	Position    string `json:"position" gorm:"column:position"`
	Height      uint   `json:"height" gorm:"column:height"`
	Weight      uint   `json:"weight" gorm:"column:weight"`
	BaseModel
}
