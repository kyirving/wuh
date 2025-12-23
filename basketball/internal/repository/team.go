package repository

import (
	"basketball/internal/model"
	"context"

	"gorm.io/gorm"
)

type TeamRepoInterface interface {
	Submit(ctx context.Context, team *model.Team) error
	Exist(ctx context.Context, team *model.Team) (bool, error)
}

type TeamRepository struct {
	Db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{Db: db}
}

func (obj *TeamRepository) Submit(ctx context.Context, team *model.Team) error {
	if err := obj.Db.WithContext(ctx).Create(team).Error; err != nil {
		return err
	}
	return nil
}

func (obj *TeamRepository) Exist(ctx context.Context, team *model.Team) (bool, error) {
	var count int64
	if err := obj.Db.WithContext(ctx).Model(&model.Team{}).Where("name = ? AND league_id = ?", team.Name, team.LeagueID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
