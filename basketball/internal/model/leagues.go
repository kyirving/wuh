package model

import "time"

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

func (obj *League) AfterFind() error {
	// 转换为本地时间
	obj.CreatedAt = obj.CreatedAt.In(time.Local)
	return nil
}
