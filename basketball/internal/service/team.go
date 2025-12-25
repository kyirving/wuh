package service

import (
	"basketball/internal/dao/admin"
	"basketball/internal/model"
	"basketball/internal/repository"
	"context"
	"errors"
)

type TeamService struct {
	TeamRepo   repository.TeamRepoInterface
	LeagueRepo repository.LeagueRepoInterface
}

func NewTeamService(teamRepo repository.TeamRepoInterface, leagueRepo repository.LeagueRepoInterface) *TeamService {
	return &TeamService{
		TeamRepo:   teamRepo,
		LeagueRepo: leagueRepo,
	}
}

// Submit 提交球队数据，校验联赛存在与球队重复
func (obj *TeamService) Submit(ctx context.Context, team *model.Team) error {
	leagueExist, err := obj.LeagueRepo.FindByID(ctx, uint(team.LeagueID))
	if err != nil {
		return err
	}
	if leagueExist == nil {
		return errors.New("league not exist")
	}

	if team.ID == 0 {
		exist, err := obj.TeamRepo.Exist(ctx, team)
		if err != nil {
			return err
		}
		if exist {
			return errors.New("team already exist")
		}
	}
	return obj.TeamRepo.Submit(ctx, team)
}

// List 查询球队列表，包含所属联赛信息
func (obj *TeamService) List(ctx context.Context) ([]admin.TeamListResp, error) {
	resp, err := obj.TeamRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	var teamList []admin.TeamListResp
	for _, team := range resp {
		teamList = append(teamList, admin.TeamListResp{
			ID:          int64(team.ID),
			Name:        team.Name,
			Description: team.Description,
			LeagueID:    int64(team.LeagueID),
			LeagueName:  team.League.Name,
			Status:      int64(team.Status),
			CreateTime:  team.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateTime:  team.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return teamList, nil
}
