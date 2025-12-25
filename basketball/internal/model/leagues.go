package model

type League struct {
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	Status      int64  `json:"status" gorm:"column:status"`
	BaseModel   `gorm:"embedded"`
}

func (League) TableName() string {
	return "leagues"
}

func NewLeague(name, description string) *League {
	return &League{
		Name:        name,
		Description: description,
		Status:      1,
	}
}
