package repository

import (
	"basketball/internal/model"
	"context"

	"gorm.io/gorm"
)

type PlayerRepositoryInterface interface {
	Create(ctx context.Context, player *model.Player) error
}

type PlayerRepo struct {
	Db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *PlayerRepo {
	return &PlayerRepo{Db: db}
}

func (obj *PlayerRepo) Create(ctx context.Context, player *model.Player) error {
	return obj.Db.WithContext(ctx).Create(player).Error
}
