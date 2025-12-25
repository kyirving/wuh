package repository

import (
	"basketball/internal/dao/admin"
	"basketball/internal/model"
	"context"

	"gorm.io/gorm"
)

type PlayerRepositoryInterface interface {
	Create(ctx context.Context, player *model.Player) error
	List(ctx context.Context, req admin.PlayerListReq) ([]model.Player, int64, error)
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

func (obj *PlayerRepo) List(ctx context.Context, req admin.PlayerListReq) ([]model.Player, int64, error) {

	query := obj.Db.WithContext(ctx).Model(&model.Player{})
	var players []model.Player
	var total int64

	if req.Name != "" {
		query = query.Where("name LIKE ?", "%"+req.Name+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.
		Preload("Team", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID, Name")
		}).
		Find(&players).Error
	if err != nil {
		return nil, 0, err
	}

	return players, total, nil
}
