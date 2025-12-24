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
	FindByID(ctx context.Context, id uint) (*model.League, error)
	List(ctx context.Context) ([]*model.League, error)
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

func (obj *LeagueRepo) FindByID(ctx context.Context, id uint) (*model.League, error) {
	var league model.League
	err := obj.Db.WithContext(ctx).First(&league, id).Error
	if err != nil {
		return nil, err
	}
	return &league, nil
}

func (obj *LeagueRepo) List(ctx context.Context) ([]*model.League, error) {
	var leagues []*model.League
	err := obj.Db.WithContext(ctx).Find(&leagues).Error
	if err != nil {
		return nil, err
	}
	return leagues, nil
}
