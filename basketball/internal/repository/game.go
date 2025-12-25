package repository

import (
	"basketball/internal/dao/admin"
	"basketball/internal/model"
	"context"

	"gorm.io/gorm"
)

type GameRepository struct {
	Db *gorm.DB
}

type GameRepositoryInterface interface {
	Create(ctx context.Context, game *model.Game) error
	List(ctx context.Context, req admin.GameListReq) ([]model.Game, int64, error)
	Update(ctx context.Context, game *model.Game) error
}

func NewGameRepository(db *gorm.DB) *GameRepository {
	return &GameRepository{Db: db}
}

func (r *GameRepository) Create(ctx context.Context, game *model.Game) error {
	return r.Db.WithContext(ctx).Create(game).Error
}

func (r *GameRepository) Update(ctx context.Context, game *model.Game) error {
	return r.Db.WithContext(ctx).Save(game).Error
}

func (r *GameRepository) List(ctx context.Context, req admin.GameListReq) ([]model.Game, int64, error) {
	var games []model.Game
	var total int64
	if err := r.Db.WithContext(ctx).Model(&model.Game{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.Db.WithContext(ctx).Model(&model.Game{}).Limit(int(req.PageInfo.PageSize)).Offset(int(req.PageInfo.PageSize * (req.PageInfo.Page - 1))).Find(&games).Error; err != nil {
		return nil, 0, err
	}
	return games, total, nil
}
