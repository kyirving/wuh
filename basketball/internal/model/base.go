package model

import (
	"time"
)

type BaseModel struct {
	ID        int       `json:"id" gorm:"primarykey"`
	Status    int       `json:"status" gorm:"default:1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
