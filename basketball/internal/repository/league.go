package repository

import (
	"basketball/internal/model"
	"context"

	"gorm.io/gorm"
)

type LeagueRepo struct {
	Db *gorm.DB
}

type LeagueRepoInterface interface {
	Submit(ctx context.Context, league *model.League) error
	Exist(ctx context.Context, name string) (bool, error)
}

func NewLeagueRepo(db *gorm.DB) *LeagueRepo {
	return &LeagueRepo{Db: db}
}

func (obj *LeagueRepo) Submit(ctx context.Context, league *model.League) error {
	return obj.Db.WithContext(ctx).Create(league).Error
}

func (obj *LeagueRepo) Exist(ctx context.Context, name string) (bool, error) {
	var count int64
	err := obj.Db.WithContext(ctx).Model(&model.League{}).Where("name = ?", name).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
