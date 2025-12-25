package model

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LeagueID    int64  `json:"league_id"`
	League      League `json:"league" gorm:"foreignKey:LeagueID;references:ID"` //foreignKey 重写外键字段  references 引用主键字段
	Status      int64  `json:"status"`
	BaseModel
}

func (Team) TableName() string {
	return "teams"
}

func (t *Team) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now()
	return nil
}
