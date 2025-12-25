package repository

import (
	"basketball/internal/model"
	"context"

	"gorm.io/gorm"
)

type TeamRepoInterface interface {
	Submit(ctx context.Context, team *model.Team) error
	Exist(ctx context.Context, team *model.Team) (bool, error)
	List(ctx context.Context) ([]*model.Team, error)
}

type TeamRepository struct {
	Db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{Db: db}
}

func (obj *TeamRepository) Submit(ctx context.Context, team *model.Team) error {
	if team.ID == 0 {
		if err := obj.Db.WithContext(ctx).Create(team).Error; err != nil {
			return err
		}
		return nil
	}
	if err := obj.Db.WithContext(ctx).Model(team).Updates(map[string]interface{}{
		"name":        team.Name,
		"description": team.Description,
		"league_id":   team.LeagueID,
		"status":      team.Status,
	}).Error; err != nil {
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

// List 返回球队列表并预加载所属联赛信息
func (obj *TeamRepository) List(ctx context.Context) ([]*model.Team, error) {
	var teams []*model.Team
	if err := obj.Db.WithContext(ctx).Model(&model.Team{}).Preload("League", func(db *gorm.DB) *gorm.DB {
		// 只选择联赛的ID和名称
		return db.Select("ID, Name")
	}).Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}
